package controllers

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/api/constants"
	"phl-skate-sharpening-api/core/models"
	"phl-skate-sharpening-api/utils"
)

type ModelController struct {
	Log *logrus.Logger
}

type ModelResponse struct {
	Models []ModelInfoStruct `json:"Models"`
}

type ModelInfoStruct struct {
	ID    int             `json:"id"`
	Name  string          `json:"name"`
	Alias string          `json:"alias"`
	Brand BrandInfoStruct `json:"brand"`
}

func (controller ModelController) GetModels(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::ModelController::GetModels - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	models, err := context.Core.ModelService.GetAllModels()
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::ModelController::GetModels - failed to get models",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructModelResponse(models)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func (controller ModelController) GetModelById(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_PARAMS, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::ModelController::GetModelById - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	modelId, err := strconv.Atoi(context.Params.ByName("modelId"))
	if err != nil {
		http.Error(writer, `{"error": "No valid id was provided"}`, http.StatusBadRequest)
		return
	}

	model, err := context.Core.ModelService.GetModelById(modelId)
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructModelResponse(model)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func ConstructModelResponse(models []models.ModelStruct) ModelResponse {
	var response ModelResponse
	modelsArray := make([]ModelInfoStruct, 0)
	for _, model := range models {
		modelResponse := ConstructModelStructResponse(model)
		modelsArray = append(modelsArray, modelResponse)
	}

	response.Models = modelsArray
	return response
}

func ConstructModelStructResponse(model models.ModelStruct) ModelInfoStruct {
	return ModelInfoStruct{
		ID:    model.ID,
		Name:  model.Name,
		Alias: model.Alias,
		Brand: ConstructBrandStructResponse(model.Brand),
	}
}
