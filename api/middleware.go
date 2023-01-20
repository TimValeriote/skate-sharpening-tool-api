package main

import (
	"database/sql"
	"net/http"
	"runtime/debug"

	"github.com/Tomasen/realip"
	"github.com/gorilla/context"
	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/api/constants"
	"phl-skate-sharpening-api/core"
	"phl-skate-sharpening-api/core/models"
	"phl-skate-sharpening-api/utils"
)

type Middleware struct {
	Database *sql.DB
	Log      *logrus.Logger
}

// Make OPTIONS pre-flight requests not die a horrible death.
func (middleware *Middleware) OptionsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
		writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Forwarded-For, X-Requested-With, Content-Type, Accept, Authorization, UUID, STREAMING_AUTH_TOKEN")
		writer.Header().Set("Content-Type", "application/json")

		// Set the correct remote address as it could be a forwarded address through a proxy
		request.RemoteAddr = realip.FromRequest(request)

		handler.ServeHTTP(writer, request)
	})
}

// Sets a bunch of headers so the API can be called from anywhere.  It's probably a good idea to restrict the origin
// to a known site eventually.
func (middleware *Middleware) CORSMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, HEAD, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Forwarded-For, X-Requested-With, Content-Type, Accept, Authorization, UUID, STREAMING_AUTH_TOKEN")
		writer.Header().Set("Content-Type", "application/json")

		// Set the correct remote address as it could be a forwarded address through a proxy
		request.RemoteAddr = realip.FromRequest(request)

		handler.ServeHTTP(writer, request)
	})
}

func (middleware *Middleware) CoreMasterCoreMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		context.Set(request, constants.CONTEXT_USEDATABASE, constants.CONTEXT_DATABASE)
		handler.ServeHTTP(writer, request)
	})
}

func (middleware *Middleware) CoreApplicationServiceMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var db *sql.DB

		service, err := utils.NewServiceFromContext(request, constants.CONTEXT_USEDATABASE)
		if err != nil {
			service.Log.WithFields(logrus.Fields{
				"event":      "phlapi::CoreApplicationServiceMiddleware - Failed to get value from context",
				"stackTrace": string(debug.Stack()),
			}).Error(err)

			http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
			return
		}

		db = middleware.Database

		co, err := core.CreateCore(db)
		if err != nil {
			service.Log.WithFields(logrus.Fields{
				"event":      "phlapi::CoreApplicationServiceMiddleware - Failed to create BACore instance",
				"stackTrace": string(debug.Stack()),
			}).Error(err)

			http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
			return
		}

		co.SetLoggerEntry(service.Log)
		co.Begin()
		defer rollbackBACoreApplicationService(co, service.Log)

		context.Set(request, constants.CONTEXT_CORE, co)
		handler.ServeHTTP(writer, request)
	})
}

func rollbackBACoreApplicationService(co *models.Core, log *logrus.Entry) {
	err := co.Rollback()
	if err != nil {
		log.WithFields(logrus.Fields{
			"event":      "V2::rollbackBACoreApplicationService - Failed when trying to rollback BACore",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
	}
}

func InitMiddleware(database *sql.DB, log *logrus.Logger) Middleware {
	var middleware Middleware
	middleware.Database = database
	middleware.Log = log
	return middleware
}
