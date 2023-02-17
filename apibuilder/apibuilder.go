package apibuilder

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/sirupsen/logrus"
)

type ApiStruct struct {
	Router                  *httprouter.Router
	Prefix                  string
	Routes                  []Route
	ConstantMiddlewareChain alice.Chain
	OptionsMiddlewareChain  alice.Chain
	Log                     *logrus.Logger
}

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware alice.Chain
}

func NewApi(prefix string, router *httprouter.Router, constantMiddlewareChain alice.Chain, optionsMiddlewareChain alice.Chain, log *logrus.Logger) *ApiStruct {
	newApi := &ApiStruct{
		Router:                  router,
		Prefix:                  prefix,
		ConstantMiddlewareChain: constantMiddlewareChain,
		OptionsMiddlewareChain:  optionsMiddlewareChain,
		Log:                     log,
	}
	return newApi
}

func (store *ApiStruct) Finalize() *ApiStruct {

	// Setup all of the routes
	uniquePaths := make(map[string]bool)
	for i := range store.Routes {
		uniquePaths[fmt.Sprintf("%s%s", store.Prefix, store.Routes[i].Path)] = true
		store.Router.Handle(store.Routes[i].Method, fmt.Sprintf("%s%s", store.Prefix, store.Routes[i].Path), store.httpWrapper(store.ConstantMiddlewareChain.Extend(store.Routes[i].Middleware).ThenFunc(store.Routes[i].Handler)))
	}

	// Setup OPTIONS request routes
	for route := range uniquePaths {
		store.Router.Handle(http.MethodOptions, route, store.httpWrapper(store.OptionsMiddlewareChain.ThenFunc(optionsController)))
	}

	// A must have to make sure the server doesn't crash if something goes wrong at any point.
	store.Router.PanicHandler = store.PanicHandler

	return store
}

func optionsController(writer http.ResponseWriter, request *http.Request) {
	//Yes, empty.  This is only here to handle pre-flight OPTIONS requests.
}

func (store *ApiStruct) httpWrapper(handler http.Handler) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// Get and set the url query params in context if they are valid
		context.Set(request, "params", params)

		handler.ServeHTTP(writer, request)
	}
}

//
func (store *ApiStruct) PanicHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")

	store.Log.WithFields(logrus.Fields{
		"event":      "apisetup::PanicHandler - Something went terribly wrong",
		"apiPrefix":  store.Prefix,
		"url":        request.URL.String(),
		"stacktrace": string(debug.Stack()),
	}).Error(err)

	var responseCode int
	if err == http.StatusNotFound {
		responseCode = http.StatusNotFound
	} else {
		responseCode = http.StatusInternalServerError
	}

	message := http.StatusText(responseCode)
	http.Error(writer, fmt.Sprintf(`{"error": "%s"}`, message), responseCode)
}
