package main

import (
	"database/sql"

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
		//This is the index route, really only for testing and to confirm the API is running on your local enviro
		{"GET", "/index", controllers.IndexController{Log: log}.Index, alice.New()},

		{"GET", "/users", controllers.UserController{Log: log}.GetUsers, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"GET", "/user/byid/:userId", controllers.UserController{Log: log}.GetUserById, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"GET", "/user/byemail", controllers.UserController{Log: log}.GetUserByEmail, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"POST", "/users/new", controllers.UserController{Log: log}.CreateUser, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"PUT", "/users/update/:userId", controllers.UserController{Log: log}.UpdateUser, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},

		{"GET", "/colours", controllers.ColourController{Log: log}.GetColours, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"GET", "/colours/:colourId", controllers.ColourController{Log: log}.GetColourById, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},

		{"GET", "/userskates/:userId", controllers.UserSkateController{Log: log}.GetAllUserSkatesByUserID, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"GET", "/userskates/:userId/:userSkateId", controllers.UserSkateController{Log: log}.GetUserSkateByUserIdAndUserSkateId, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"GET", "/userskatesnotsharpening/:userId", controllers.UserSkateController{Log: log}.GetAllUserSkatesNotBeingSharpenedByUserId, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"POST", "/userskate/new/:userId", controllers.UserSkateController{Log: log}.CreateUserSkate, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"PUT", "/userskate/update/:userSkateId", controllers.UserSkateController{Log: log}.UpdateUserSkate, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"DELETE", "/userskate/delete/:userId/:userSkateId", controllers.UserSkateController{Log: log}.DeleteUserSkate, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},

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

		{"GET", "/sharpenings/open/:userId", controllers.SharpeningController{Log: log}.GetOpenSharpeningsForUser, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"DELETE", "/sharpenings/delete/open/:sharpenId/:userId", controllers.SharpeningController{Log: log}.DeleteSharpen, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
		{"POST", "/sharpenings/create/:userId/:userSkateId/:storeId", controllers.SharpeningController{Log: log}.CreateNewUserSharpening, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},

		{"GET", "/sharpeningcodes/:code", controllers.SharpeningCodeController{Log: log}.CheckIfCodeIsValid, alice.New(middleware.CoreMasterCoreMiddleware, middleware.CoreApplicationServiceMiddleware)},
	}

	api.Finalize()
}
