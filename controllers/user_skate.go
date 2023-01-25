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

type UserSkateController struct {
	Log *logrus.Logger
}

type UserSkateResponse struct {
	Skates []UserSkatesInfoStruct `json:"Skates"`
}

type UserSkatesInfoStruct struct {
	ID              int              `json:"id"`
	Skate           SkateInfoStruct  `json:"skate"`
	Holder          BrandInfoStruct  `json:"holder"`
	HolderSize      float32          `json:"holder_size"`
	SkateSize       float32          `json:"skate_size"`
	LaceColour      ColourInfoStruct `json:"lace_colour"`
	HasSteel        bool             `json:"has_steel"`
	Steel           BrandInfoStruct  `json:"steel"`
	HasGuards       bool             `json:"has_guards"`
	GuardColour     ColourInfoStruct `json:"guard_colour"`
	PreferredRadius string           `json:"preferred_radius"`
	Fit             FitInfoStruct    `json:"fit"`
}

func (controller UserSkateController) GetAllUserSkatesByUserID(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_PARAMS, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::UserSkateController::GetAllUserSkatesByUserID - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	userId, err := strconv.Atoi(context.Params.ByName("userId"))
	if err != nil {
		http.Error(writer, `{"error": "No valid id was provided"}`, http.StatusBadRequest)
		return
	}

	userSkates, err := context.Core.UserSkateService.GetAllUserSkatesByUserID(userId)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::UserSkateController::GetAllUserSkatesByUserID - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructUserSkatesResponse(userSkates)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func (controller UserSkateController) GetUserSkateByUserIdAndUserSkateId(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_PARAMS, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::UserSkateController::GetUserSkateByUserIdAndUserSkateId - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	userId, err := strconv.Atoi(context.Params.ByName("userId"))
	if err != nil {
		http.Error(writer, `{"error": "No valid id was provided"}`, http.StatusBadRequest)
		return
	}

	userSkateId, err := strconv.Atoi(context.Params.ByName("userSkateId"))
	if err != nil {
		http.Error(writer, `{"error": "No valid id was provided"}`, http.StatusBadRequest)
		return
	}

	userSkates, err := context.Core.UserSkateService.GetUserSkateByUserIdAndUserSkateId(userId, userSkateId)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::UserSkateController::GetUserSkateByUserIdAndUserSkateId - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructUserSkatesResponse(userSkates)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func ConstructUserSkatesResponse(userSkates []models.UserSkateStruct) UserSkateResponse {
	var response UserSkateResponse
	skatesArray := make([]UserSkatesInfoStruct, 0)
	for _, userSkate := range userSkates {
		var userReponse UserSkatesInfoStruct
		userReponse.ID = userSkate.ID
		userReponse.Skate = ConstructSkatesResponse(userSkate.Skate)
		userReponse.Holder = ConstructBrandStructResponse(userSkate.Holder)
		userReponse.HolderSize = userSkate.HolderSize
		userReponse.SkateSize = userSkate.SkateSize
		userReponse.LaceColour = ConstructColourStructResponse(userSkate.LaceColour)
		userReponse.HasSteel = userSkate.HasSteel
		userReponse.HasGuards = userSkate.HasGuards
		userReponse.GuardColour = ConstructColourStructResponse(userSkate.GuardColour)
		userReponse.Fit = ConstructFitStructResponse(userSkate.Fit)
		userReponse.PreferredRadius = userSkate.PreferredRadius
		skatesArray = append(skatesArray, userReponse)
	}

	response.Skates = skatesArray
	return response
}

func ConstructSkatesResponse(skate models.SkateStruct) SkateInfoStruct {
	return SkateInfoStruct{
		ID:    skate.ID,
		Model: ConstructModelStructResponse(skate.Model),
		Brand: ConstructBrandStructResponse(skate.Brand),
	}
}
