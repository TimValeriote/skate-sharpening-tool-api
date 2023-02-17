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

type ColourController struct {
	Log *logrus.Logger
}

type ColoursResponse struct {
	Colours []ColourInfoStruct `json:"Colours"`
}

type ColourInfoStruct struct {
	ID     int    `json:"id"`
	Colour string `json:"colour"`
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

func (controller ColourController) GetColourById(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_PARAMS, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::ColourController::GetColourById - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	colourId, err := strconv.Atoi(context.Params.ByName("colourId"))
	if err != nil {
		http.Error(writer, `{"error": "No valid id was provided"}`, http.StatusBadRequest)
		return
	}

	colour, err := context.Core.ColourService.GetColourById(colourId)
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructColourResponse(colour)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func ConstructColourResponse(colours []models.ColourStruct) ColoursResponse {
	var response ColoursResponse
	coloursArray := make([]ColourInfoStruct, 0)
	for _, colour := range colours {
		var colourReponse ColourInfoStruct
		colourReponse.ID = colour.ID
		colourReponse.Colour = colour.Colour
		coloursArray = append(coloursArray, colourReponse)
	}

	response.Colours = coloursArray
	return response
}

func ConstructColourStructResponse(colour models.ColourStruct) ColourInfoStruct {
	var colourReponse ColourInfoStruct
	colourReponse.ID = colour.ID
	colourReponse.Colour = colour.Colour
	return colourReponse
}

func ConstructColourStructResponseBasedOnBool(colour models.ColourStruct, isColour bool) ColourInfoStruct {
	var colourReponse ColourInfoStruct
	if !isColour {
		colourReponse.ID = 0
		colourReponse.Colour = "null"
	} else {
		colourReponse.ID = colour.ID
		colourReponse.Colour = colour.Colour
	}
	return colourReponse
}
