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

type PeopleController struct {
	Log *logrus.Logger
}

type PeopleResponse struct {
	People []PeopleInfoStruct `json:"People"`
}

type PeopleInfoStruct struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (controller PeopleController) GetPeople(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::PeopleController::GetPeople - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	people, err := context.Core.PeopleService.GetAllPeople()
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructPeopleResponse(people)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func (controller PeopleController) CreatePerson(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_PARAMS, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::PeopleController::GetPeople - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	var newPerson models.PeopleStruct
	newPerson.FirstName = request.URL.Query().Get("first_name")
	newPerson.LastName = request.URL.Query().Get("last_name")

	if len(newPerson.FirstName) <= 0 {
		http.Error(writer, `{"error": "First name is required"}`, http.StatusInternalServerError)
		return
	} else if len(newPerson.LastName) <= 0 {
		http.Error(writer, `{"error": "Last name is required"}`, http.StatusInternalServerError)
		return
	}

	responseStatus := http.StatusCreated

	err = context.Core.PeopleService.CreatePerson(&newPerson)
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	} else {
		responseStatus = http.StatusOK
	}

	context.Core.Commit()
	writer.WriteHeader(responseStatus)
}

func ConstructPeopleResponse(people []models.PeopleStruct) PeopleResponse {
	var response PeopleResponse
	peopleArray := make([]PeopleInfoStruct, 0)
	for _, person := range people {
		var personInfo PeopleInfoStruct
		personInfo.FirstName = person.FirstName
		personInfo.LastName = person.LastName
		peopleArray = append(peopleArray, personInfo)
	}

	response.People = peopleArray
	return response
}
