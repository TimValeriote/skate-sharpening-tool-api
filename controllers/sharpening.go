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

type SharpeningController struct {
	Log *logrus.Logger
}

type SharpeningResponse struct {
	Sharpenings []SharpeningInfoStruct `json:"Sharpenings"`
}

type SharpeningInfoStruct struct {
	ID          int `json:"id"`
	UserId      int `json:"user_id"`
	UserSkateId int `json:"user_skate_id"`
	StoreId     int `json:"store_id"`
}

func (controller SharpeningController) GetOpenSharpeningsForUser(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_PARAMS, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::SharpeningController::GetOpenSharpeningsForUser - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	userId, err := strconv.Atoi(context.Params.ByName("userId"))
	if err != nil {
		http.Error(writer, `{"error": "No valid user id was provided"}`, http.StatusBadRequest)
		return
	}

	openSharpenings, err := context.Core.SharpeningService.GetOpenSharpeningsForUser(userId)
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructUserSharpeningInfoDetailsStruct(openSharpenings)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func ConstructUserSharpeningInfoDetailsStruct(sharpenings []models.SharpeningStruct) SharpeningResponse {
	var response SharpeningResponse
	sharpeningsArray := make([]SharpeningInfoStruct, 0)
	for _, sharpening := range sharpenings {
		var indivSharpeningResp SharpeningInfoStruct
		indivSharpeningResp.ID = sharpening.ID
		indivSharpeningResp.UserId = sharpening.UserId
		indivSharpeningResp.UserSkateId = sharpening.UserSkateId
		indivSharpeningResp.StoreId = sharpening.StoreId
		sharpeningsArray = append(sharpeningsArray, indivSharpeningResp)
	}
	response.Sharpenings = sharpeningsArray
	return response
}

/*func ConstructSharpeningInfoResponse(sharpenings []models.SharpeningStruct) SharpeningResponse {
	var response SharpeningResponse
	usersArray := make([]SharpeningInfoStruct, 0)
	for _, sharpening := range sharpenings {
		var sharpeningResponse SharpeningInfoStruct
		sharpeningResponse.ID = sharpening.ID
		sharpeningResponse.UserId = sharpening.FirstName
		sharpeningResponse.LastName = sharpening.LastName
		sharpeningResponse.Email = sharpening.Email
		usersArray = append(usersArray, sharpeningResponse)
	}

	response.Users = usersArray
	return response
}

func ConstructUserInfoResponse(user models.UsersStruct) UserInfoStruct {
	return UserInfoStruct{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		UUID:        user.UUID,
		IsStaff:     user.IsStaff,
	}
}

func validateNewUser(user models.UsersStruct) (bool, string) {
	if len(user.FirstName) <= 0 {
		return false, "first_name is missing"
	} else if len(user.LastName) <= 0 {
		return false, "last_name is missing"
	} else if len(user.Email) <= 0 {
		return false, "email is missing"
	} else if len(user.PhoneNumber) <= 0 {
		return false, "phone_number is missing"
	}
	return true, ""
}

func validateUpdatedUser(user models.UpdateUserStruct) (bool, string) {
	if len(user.FirstName) <= 0 {
		return false, "first_name is missing"
	} else if len(user.LastName) <= 0 {
		return false, "last_name is missing"
	} else if len(user.Email) <= 0 {
		return false, "email is missing"
	} else if len(user.PhoneNumber) <= 0 {
		return false, "phone_number is missing"
	}
	return true, ""
}*/
