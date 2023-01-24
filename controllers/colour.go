package controllers

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"runtime/debug"

	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/api/constants"
	"phl-skate-sharpening-api/core/models"
	"phl-skate-sharpening-api/utils"
)

type ColourController struct {
	Log *logrus.Logger
}

type ColoursResponse struct {
	Colours []ColourInfoStruct `json:"Colours"`
}

type ColourInfoStruct struct {
	ID     int    `json:"id"`
	Colour string `json:"Colour"`
}

func (controller ColourController) GetColours(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::ColourController::GetUsers - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	colours, err := context.Core.ColourService.GetAllColours()
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructColourResponse(colours)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func (controller ColourController) GetColourByName(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_PARAMS, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::ColourController::GetPerson - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	colourName := request.URL.Query().Get("colourName")
	if len(colourName) <= 0 {
		http.Error(writer, `{"error": "Colour name is required"}`, http.StatusInternalServerError)
		return
	}

	user, err := context.Core.ColourService.GetColourByName(colourName)
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructColourResponse(user)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func ConstructColourResponse(colours []models.ColourStruct) ColoursResponse {
	var response ColoursResponse
	coloursArray := make([]ColourInfoStruct, 0)
	for _, colour := range colours {
		var colourReponse ColourInfoStruct
		if !colour.ID.Valid {
			colourReponse.ID = 0
		}
		colourReponse.Colour = colour.Colour.ValueOrZero()
		coloursArray = append(coloursArray, colourReponse)
	}

	response.Colours = coloursArray
	return response
}

func ConstructColourStructResponse(colour models.ColourStruct) ColourInfoStruct {
	var colourReponse ColourInfoStruct
	if !colour.ID.Valid {
		colourReponse.ID = 0
	}
	colourReponse.Colour = colour.Colour.ValueOrZero()
	return colourReponse
}
