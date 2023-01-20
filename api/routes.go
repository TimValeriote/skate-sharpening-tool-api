package main

import (
	"database/sql"
	//"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/apibuilder"
	"phl-skate-sharpening-api/controllers"
)

func SetupRouting(router *httprouter.Router, database *sql.DB, log *logrus.Logger) {
	middleware := InitMiddleware(database, log)
	api := apibuilder.NewApi("/phlapi", router, alice.New(context.ClearHandler, middleware.CORSMiddleware), alice.New(context.ClearHandler, middleware.OptionsMiddleware), log)

	api.Routes = []apibuilder.Route{
		{"GET", "/test", controllers.IndexController{Log: log}.Index, alice.New()},

		{"GET", "/people", controllers.PeopleController{Log: log}.GetPeople, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"GET", "/people/new", controllers.PeopleController{Log: log}.CreatePerson, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
	}

	api.Finalize()
}
