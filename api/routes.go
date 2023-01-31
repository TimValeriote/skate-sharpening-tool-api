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
	api := apibuilder.NewApi("/sstapi", router, alice.New(context.ClearHandler, middleware.CORSMiddleware), alice.New(context.ClearHandler, middleware.OptionsMiddleware), log)

	api.Routes = []apibuilder.Route{
		{"GET", "/index", controllers.IndexController{Log: log}.Index, alice.New()},

		//These are examples, there is no data in the db containing people
		{"GET", "/people", controllers.PeopleController{Log: log}.GetPeople, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"GET", "/person/:personId", controllers.PeopleController{Log: log}.GetPerson, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"POST", "/people/new", controllers.PeopleController{Log: log}.CreatePerson, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"PUT", "/people/update/:personId", controllers.PeopleController{Log: log}.UpdatePersonByID, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},

		{"GET", "/users", controllers.UserController{Log: log}.GetUsers, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"GET", "/user/byid/:userId", controllers.UserController{Log: log}.GetUserById, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"GET", "/user/byemail", controllers.UserController{Log: log}.GetUserByEmail, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"POST", "/users/new", controllers.UserController{Log: log}.CreateUser, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"PUT", "/users/update/:userId", controllers.UserController{Log: log}.UpdateUser, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},

		{"GET", "/colours", controllers.ColourController{Log: log}.GetColours, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"GET", "/colours/:colourId", controllers.ColourController{Log: log}.GetColourById, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},

		{"GET", "/userskates/:userId", controllers.UserSkateController{Log: log}.GetAllUserSkatesByUserID, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"GET", "/userskates/:userId/:userSkateId", controllers.UserSkateController{Log: log}.GetUserSkateByUserIdAndUserSkateId, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"POST", "/userskate/new/:userId", controllers.UserSkateController{Log: log}.CreateUserSkate, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"PUT", "/userskate/update/:userSkateId", controllers.UserSkateController{Log: log}.UpdateUserSkate, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},

		{"GET", "/fits", controllers.FitController{Log: log}.GetFits, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"GET", "/fits/:fitId", controllers.FitController{Log: log}.GetFitsById, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},

		{"GET", "/brands", controllers.BrandController{Log: log}.GetBrands, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"GET", "/brands/:brandId", controllers.BrandController{Log: log}.GetBrandById, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},

		{"GET", "/models", controllers.ModelController{Log: log}.GetModels, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"GET", "/models/:modelId", controllers.ModelController{Log: log}.GetModelById, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},

		{"GET", "/skates", controllers.SkateController{Log: log}.GetSkates, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"GET", "/skates/:skateId", controllers.SkateController{Log: log}.GetSkateById, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},

		{"GET", "/stores", controllers.StoreController{Log: log}.GetStores, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"GET", "/stores/:storeId", controllers.StoreController{Log: log}.GetStoreById, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
	}

	api.Finalize()
}
