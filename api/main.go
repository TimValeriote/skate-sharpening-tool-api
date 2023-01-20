package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"runtime/debug"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/config"
	"phl-skate-sharpening-api/core"
	"phl-skate-sharpening-api/utils"
)

var database *sql.DB
var logger *logrus.Logger

func main() {
	var err error
	var configFile string

	logger = logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	flag.StringVar(&configFile, "with-config", "", "Path to the config file")
	flag.Parse()

	if len(configFile) == 0 {
		logger.Fatal("Must set --with-config option")
	}

	logger.Info("Loading server configuration from config file.")

	config.Config, err = config.NewConfiguration(configFile)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"event":      "Error loading config file",
			"stacktrace": string(debug.Stack()),
		}).Fatal(err)
	}

	err = config.Config.Validate()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"event":      "Error in server configuration",
			"stacktrace": string(debug.Stack()),
		}).Fatal(err)
	} else {
		logger.Info("Server configuration is valid.")
	}

	database, err = core.ConnectToDatabase(
		config.Config.Database.Server,
		config.Config.Database.Port,
		config.Config.Database.Username,
		config.Config.Database.Password,
		config.Config.Database.Name)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"event":      "Could not connect to the database",
			"stacktrace": string(debug.Stack()),
		}).Fatal(err)
	}

	router := httprouter.New()

	SetupRouting(router, database, logger)

	go utils.LogsHeartbeat(logger, 20)

	logger.Infof("Server is starting up on port %d.", config.Config.ListenPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Config.ListenPort), router)
}
