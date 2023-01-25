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

type UserController struct {
	Log *logrus.Logger
}

type UserResponse struct {
	Users []UserInfoStruct `json:"Users"`
}

type UserInfoStruct struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	UUID        string `json:"uuid"`
	IsStaff     bool   `json:"is_staff"`
}

func (controller UserController) GetUsers(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::UserController::GetUsers - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	users, err := context.Core.UserService.GetAllUsers()
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructUsersResponse(users)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func (controller UserController) GetUserById(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_PARAMS, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::UserController::GetUserById - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	userId, err := strconv.Atoi(context.Params.ByName("userId"))
	if err != nil {
		http.Error(writer, `{"error": "No valid id was provided"}`, http.StatusBadRequest)
		return
	}

	user, err := context.Core.UserService.GetUserById(userId)
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructUsersResponse(user)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func ConstructUsersResponse(users []models.UsersStruct) UserResponse {
	var response UserResponse
	usersArray := make([]UserInfoStruct, 0)
	for _, user := range users {
		var userReponse UserInfoStruct
		userReponse.ID = user.ID
		userReponse.FirstName = user.FirstName
		userReponse.LastName = user.LastName
		userReponse.Email = user.Email
		userReponse.PhoneNumber = user.PhoneNumber
		userReponse.UUID = user.UUID
		userReponse.IsStaff = user.IsStaff
		usersArray = append(usersArray, userReponse)
	}

	response.Users = usersArray
	return response
}
