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

type StoreController struct {
	Log *logrus.Logger
}

type StoreResponse struct {
	Stores []StoreInfoStruct `json:"Stores"`
}

type StoreInfoStruct struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phone_number"`
	StoreNumber string `json:"store_number"`
}

func (controller StoreController) GetStores(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::StoreController::GetStores - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	stores, err := context.Core.StoreService.GetAllStores()
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructStoreResponse(stores)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func (controller StoreController) GetStoreById(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_PARAMS, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::StoreController::GetStoreById - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	storeId, err := strconv.Atoi(context.Params.ByName("storeId"))
	if err != nil {
		http.Error(writer, `{"error": "No valid id was provided"}`, http.StatusBadRequest)
		return
	}

	sotre, err := context.Core.StoreService.GetStoreById(storeId)
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructStoreResponse(sotre)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func ConstructStoreResponse(stores []models.StoreStruct) StoreResponse {
	var response StoreResponse
	storesArray := make([]StoreInfoStruct, 0)
	for _, store := range stores {
		storeResponse := ConstructStoreStructResponse(store)
		storesArray = append(storesArray, storeResponse)
	}

	response.Stores = storesArray
	return response
}

func ConstructStoreStructResponse(store models.StoreStruct) StoreInfoStruct {
	return StoreInfoStruct{
		ID:          store.ID,
		Name:        store.Name,
		Address:     store.Address,
		City:        store.City,
		Country:     store.Country,
		PhoneNumber: store.PhoneNumber,
		StoreNumber: store.StoreNumber,
	}
}
