package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type IndexController struct {
	Log *logrus.Logger
}

type IndexResponse struct {
	Status string `json:"status"`
}

func (controller IndexController) Index(writer http.ResponseWriter, request *http.Request) {
	controller.Log.Info("Got to the Index")

	// Generate the response before returning in-case something goes boom
	response := ConstructIndexResponse("success")

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func ConstructIndexResponse(status string) *IndexResponse {
	return &IndexResponse{
		Status: status,
	}
}
