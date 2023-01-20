package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/api/constants"
	"phl-skate-sharpening-api/core/models"
)

type ControllerContextServiceStruct struct {
	Log         *logrus.Entry
	Params      httprouter.Params
	UseDatabase string
	Core        *models.Core
}

func LogsHeartbeat(log *logrus.Logger, loggingInterval int) {
	logHeartBeatInterval := (time.Duration(loggingInterval) * time.Minute)

	for t := range time.NewTicker(time.Duration(logHeartBeatInterval)).C {
		_ = t
		log.WithFields(logrus.Fields{
			"heartbeat": "At first I was afraid I was petrified",
		}).Info()
	}
}

func NewServiceFromContext(request *http.Request, values ...string) (*ControllerContextServiceStruct, error) {
	service := &ControllerContextServiceStruct{}

	if val, ok := context.GetOk(request, constants.CONTEXT_LOGGER); ok {
		service.Log = val.(*logrus.Entry)
	} else {
		logger := logrus.New()
		logger.Formatter = &logrus.JSONFormatter{}
		service.Log = logger.WithFields(logrus.Fields{})
	}
	for i := range values {
		switch values[i] {
		case constants.CONTEXT_PARAMS:
			if val, ok := context.GetOk(request, constants.CONTEXT_PARAMS); ok {
				service.Params = val.(httprouter.Params)
			} else {
				return nil, fmt.Errorf("Failed to get Params from context")
			}
		case constants.CONTEXT_USEDATABASE:
			if val, ok := context.GetOk(request, constants.CONTEXT_USEDATABASE); ok {
				service.UseDatabase = val.(string)
			} else {
				return service, fmt.Errorf("Failed to get UseDatabase from context")
			}
		case constants.CONTEXT_CORE:
			if val, ok := context.GetOk(request, constants.CONTEXT_CORE); ok {
				service.Core = val.(*models.Core)
			} else {
				return nil, fmt.Errorf("Failed to get Core from context")
			}
		}
	}

	return service, nil

}
