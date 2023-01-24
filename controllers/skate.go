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

type SkateController struct {
	Log *logrus.Logger
}

type SkateResponse struct {
	Skates []SkateInfoStruct `json:"Skates"`
}

type SkateInfoStruct struct {
	ID    int             `json:"id"`
	Model ModelInfoStruct `json:"model"`
	Brand BrandInfoStruct `json:"brand"`
	Fit   FitInfoStruct   `json:"fit"`
}

func (controller SkateController) GetSkates(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::SkateController::GetSkates - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	skates, err := context.Core.SkateService.GetAllSkates()
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::SkateController::GetSkates - failed to get skates",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructSkateResponse(skates)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func ConstructSkateResponse(skates []models.SkateStruct) SkateResponse {
	var response SkateResponse
	skatesArray := make([]SkateInfoStruct, 0)
	for _, skate := range skates {
		skatesResponse := ConstructSkateStructResponse(skate)
		skatesArray = append(skatesArray, skatesResponse)
	}

	response.Skates = skatesArray
	return response
}

func ConstructSkateStructResponse(skate models.SkateStruct) SkateInfoStruct {
	return SkateInfoStruct{
		ID:    skate.ID,
		Model: ConstructModelStructResponse(skate.Model),
		Brand: ConstructBrandStructResponse(skate.Brand),
		Fit:   ConstructFitStructResponse(skate.Fit),
	}
}
