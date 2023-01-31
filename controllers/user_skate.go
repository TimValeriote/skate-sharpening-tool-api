package controllers

import (
	"encoding/json"
	"fmt"
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

	response := ConstructUserSkatesInfoResponse(userSkates)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func (controller UserSkateController) CreateUserSkate(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_PARAMS, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::UserSkateController::CreateUserSkate - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	userId, err := strconv.Atoi(context.Params.ByName("userId"))
	if err != nil {
		http.Error(writer, `{"error": "No valid id was provided"}`, http.StatusBadRequest)
		return
	}

	var newUserSkate models.CreateUserSkateStruct
	newUserSkate.UserID = userId
	newUserSkate.SkateID = request.URL.Query().Get("skate_id")
	newUserSkate.HolderBrandID = request.URL.Query().Get("holder_brand_id")
	newUserSkate.HolderSize = request.URL.Query().Get("holder_size")
	newUserSkate.SkateSize = request.URL.Query().Get("skate_size")
	newUserSkate.LaceColourID = request.URL.Query().Get("lace_colour_id")
	newUserSkate.HasSteel = request.URL.Query().Get("has_steel")
	newUserSkate.SteelID = request.URL.Query().Get("steel_id")
	newUserSkate.HasGuards = request.URL.Query().Get("has_guards")
	newUserSkate.GuardColourID = request.URL.Query().Get("guard_colour_id")
	newUserSkate.PreferredRadius = request.URL.Query().Get("preferred_radius")
	newUserSkate.FitID = request.URL.Query().Get("fit_id")

	canUpdate, msg := validateNewUserSkate(newUserSkate)

	if !canUpdate {
		message := fmt.Sprintf("{\"error\": \"%s\"}", msg)
		http.Error(writer, message, http.StatusInternalServerError)
		return
	}

	responseStatus := http.StatusCreated

	_, err = context.Core.UserSkateService.CreateUserSkate(&newUserSkate)
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	} else {
		responseStatus = http.StatusOK
	}

	context.Core.Commit()
	writer.WriteHeader(responseStatus)
}

func (controller UserSkateController) UpdateUserSkate(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_PARAMS, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::UserSkateController::GetPeople - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	var updatedUserSkate models.UpdateUserSkateStruct
	userSkateId, err := strconv.Atoi(context.Params.ByName("userSkateId"))
	if err != nil {
		http.Error(writer, `{"error": "No valid id was provided"}`, http.StatusBadRequest)
		return
	}
	updatedUserSkate.UserSkateID = userSkateId
	updatedUserSkate.SkateID = request.URL.Query().Get("skate_id")
	updatedUserSkate.HolderBrandID = request.URL.Query().Get("holder_brand_id")
	updatedUserSkate.HolderSize = request.URL.Query().Get("holder_size")
	updatedUserSkate.SkateSize = request.URL.Query().Get("skate_size")
	updatedUserSkate.LaceColourID = request.URL.Query().Get("lace_colour_id")
	updatedUserSkate.HasSteel = request.URL.Query().Get("has_steel")
	updatedUserSkate.SteelID = request.URL.Query().Get("steel_id")
	updatedUserSkate.HasGuards = request.URL.Query().Get("has_guards")
	updatedUserSkate.GuardColourID = request.URL.Query().Get("guard_colour_id")
	updatedUserSkate.PreferredRadius = request.URL.Query().Get("preferred_radius")
	updatedUserSkate.FitID = request.URL.Query().Get("fit_id")

	canUpdate, msg := validateUpdatedUserSkate(updatedUserSkate)

	if !canUpdate {
		message := fmt.Sprintf("{\"error\": \"%s\"}", msg)
		http.Error(writer, message, http.StatusInternalServerError)
		return
	}

	responseStatus := http.StatusCreated

	_, err = context.Core.UserSkateService.UpdateUserSkate(updatedUserSkate)
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	} else {
		responseStatus = http.StatusOK
	}

	context.Core.Commit()
	writer.WriteHeader(responseStatus)
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
		userReponse.Steel = ConstructBrandStructResponse(userSkate.Steel)
		userReponse.GuardColour = ConstructColourStructResponseBasedOnBool(userSkate.GuardColour, userSkate.HasGuards)
		userReponse.Fit = ConstructFitStructResponse(userSkate.Fit)
		userReponse.PreferredRadius = userSkate.PreferredRadius
		skatesArray = append(skatesArray, userReponse)
	}

	response.Skates = skatesArray
	return response
}

func ConstructUserSkatesInfoResponse(userSkate models.UserSkateStruct) UserSkatesInfoStruct {
	var userSkateResponse UserSkatesInfoStruct
	userSkateResponse.ID = userSkate.ID
	userSkateResponse.Skate = ConstructSkatesResponse(userSkate.Skate)
	userSkateResponse.Holder = ConstructBrandStructResponse(userSkate.Holder)
	userSkateResponse.HolderSize = userSkate.HolderSize
	userSkateResponse.SkateSize = userSkate.SkateSize
	userSkateResponse.LaceColour = ConstructColourStructResponse(userSkate.LaceColour)
	userSkateResponse.HasSteel = userSkate.HasSteel
	userSkateResponse.HasGuards = userSkate.HasGuards
	userSkateResponse.Steel = ConstructBrandStructResponse(userSkate.Steel)
	userSkateResponse.GuardColour = ConstructColourStructResponseBasedOnBool(userSkate.GuardColour, userSkate.HasGuards)
	userSkateResponse.Fit = ConstructFitStructResponse(userSkate.Fit)
	userSkateResponse.PreferredRadius = userSkate.PreferredRadius

	return userSkateResponse
}

func ConstructSkatesResponse(skate models.SkateStruct) SkateInfoStruct {
	return SkateInfoStruct{
		ID:    skate.ID,
		Model: ConstructModelStructResponse(skate.Model),
		Brand: ConstructBrandStructResponse(skate.Brand),
	}
}

func validateNewUserSkate(newUserSkate models.CreateUserSkateStruct) (bool, string) {
	if len(newUserSkate.SkateID) <= 0 {
		return false, "skate_id is missing"
	} else if len(newUserSkate.HolderBrandID) <= 0 {
		return false, "holder_brand_id is missing"
	} else if len(newUserSkate.HolderSize) <= 0 {
		return false, "holder_size is missing"
	} else if len(newUserSkate.SkateSize) <= 0 {
		return false, "skate_size is missing"
	} else if len(newUserSkate.LaceColourID) <= 0 {
		return false, "lace_colour_id is missing"
	} else if len(newUserSkate.HasSteel) <= 0 {
		return false, "has_steel is missing"
	} else if len(newUserSkate.HasGuards) <= 0 {
		return false, "has_guards is missing"
	} else if len(newUserSkate.GuardColourID) <= 0 {
		return false, "guard_colour_id is missing"
	} else if len(newUserSkate.PreferredRadius) <= 0 {
		return false, "preferred_radius is missing"
	} else if len(newUserSkate.FitID) <= 0 {
		return false, "fit_id is missing"
	}
	return true, ""
}

func validateUpdatedUserSkate(newUserSkate models.UpdateUserSkateStruct) (bool, string) {
	if len(newUserSkate.SkateID) <= 0 {
		return false, "skate_id is missing"
	} else if len(newUserSkate.HolderBrandID) <= 0 {
		return false, "holder_brand_id is missing"
	} else if len(newUserSkate.HolderSize) <= 0 {
		return false, "holder_size is missing"
	} else if len(newUserSkate.SkateSize) <= 0 {
		return false, "skate_size is missing"
	} else if len(newUserSkate.LaceColourID) <= 0 {
		return false, "lace_colour_id is missing"
	} else if len(newUserSkate.HasSteel) <= 0 {
		return false, "has_steel is missing"
	} else if len(newUserSkate.HasGuards) <= 0 {
		return false, "has_guards is missing"
	} else if len(newUserSkate.GuardColourID) <= 0 {
		return false, "guard_colour_id is missing"
	} else if len(newUserSkate.PreferredRadius) <= 0 {
		return false, "preferred_radius is missing"
	} else if len(newUserSkate.FitID) <= 0 {
		return false, "fit_id is missing"
	}
	return true, ""
}
