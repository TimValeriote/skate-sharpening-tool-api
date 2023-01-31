package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/google/uuid"
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

	response := ConstructUserInfoResponse(user)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func (controller UserController) GetUserByEmail(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_PARAMS, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::UserController::GetUserByEmail - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	userEmail := request.URL.Query().Get("user_email")
	if len(userEmail) <= 0 {
		http.Error(writer, `{"error": "No valid user email was provided"}`, http.StatusBadRequest)
		return
	}

	user, err := context.Core.UserService.GetUserByEmail(userEmail)
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructUserInfoResponse(user)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func (controller UserController) CreateUser(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_PARAMS, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::UserController::CreateUser - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	var newUser models.UsersStruct
	newUser.FirstName = request.URL.Query().Get("first_name")
	newUser.LastName = request.URL.Query().Get("last_name")
	newUser.Email = request.URL.Query().Get("email")
	newUser.PhoneNumber = request.URL.Query().Get("phone_number")

	uuid := uuid.New()
	newUser.UUID = uuid.String()

	newUser.IsStaff = false

	canCreate, msg := validateNewUser(newUser)

	if !canCreate {
		message := fmt.Sprintf("{\"error\": \"%s\"}", msg)
		http.Error(writer, message, http.StatusInternalServerError)
		return
	}

	responseStatus := http.StatusCreated

	newUserId, err := context.Core.UserService.CreateUser(&newUser)
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	} else {
		newUser.ID = newUserId
		responseStatus = http.StatusOK
	}

	response := ConstructUserInfoResponse(newUser)

	context.Core.Commit()
	writer.WriteHeader(responseStatus)
	json.NewEncoder(writer).Encode(response)
}

func (controller UserController) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_PARAMS, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::UserController::GetPeople - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	var updatedUser models.UpdateUserStruct
	userId, err := strconv.Atoi(context.Params.ByName("userId"))
	if err != nil {
		http.Error(writer, `{"error": "No valid id was provided"}`, http.StatusBadRequest)
		return
	}
	updatedUser.UserID = userId
	updatedUser.FirstName = request.URL.Query().Get("first_name")
	updatedUser.LastName = request.URL.Query().Get("last_name")
	updatedUser.Email = request.URL.Query().Get("email")
	updatedUser.PhoneNumber = request.URL.Query().Get("phone_number")

	canUpdate, msg := validateUpdatedUser(updatedUser)

	if !canUpdate {
		message := fmt.Sprintf("{\"error\": \"%s\"}", msg)
		http.Error(writer, message, http.StatusInternalServerError)
		return
	}

	responseStatus := http.StatusCreated

	_, err = context.Core.UserService.UpdateUser(updatedUser)
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	} else {
		responseStatus = http.StatusOK
	}

	context.Core.Commit()
	writer.WriteHeader(responseStatus)
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
}
