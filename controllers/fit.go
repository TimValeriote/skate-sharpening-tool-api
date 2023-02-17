package controllers

import (
	"encoding/json"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/api/constants"
	"phl-skate-sharpening-api/core/models"
	"phl-skate-sharpening-api/utils"
)

type FitController struct {
	Log *logrus.Logger
}

type FitsResponse struct {
	Fits []FitInfoStruct `json:"fits"`
}

type FitInfoStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (controller FitController) GetFits(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::FitController::GetFits - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	users, err := context.Core.FitService.GetAllFits()
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructFitResponse(users)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func (controller FitController) GetFitsById(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_PARAMS, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::FitController::GetFitsById - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	fitId, err := strconv.Atoi(context.Params.ByName("fitId"))
	if err != nil {
		http.Error(writer, `{"error": "No valid id was provided"}`, http.StatusBadRequest)
		return
	}

	fit, err := context.Core.FitService.GetFitById(fitId)
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructFitResponse(fit)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func ConstructFitResponse(fits []models.FitStruct) FitsResponse {
	var response FitsResponse
	fitsArray := make([]FitInfoStruct, 0)
	for _, fit := range fits {
		var fitReponse FitInfoStruct
		fitReponse.ID = fit.ID
		fitReponse.Name = fit.Name
		fitsArray = append(fitsArray, fitReponse)
	}

	response.Fits = fitsArray
	return response
}

func ConstructFitStructResponse(fit models.FitStruct) FitInfoStruct {
	return FitInfoStruct{
		ID:   fit.ID,
		Name: fit.Name,
	}
}
