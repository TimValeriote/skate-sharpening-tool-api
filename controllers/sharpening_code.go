package controllers

import (
	"encoding/json"
	"net/http"
	"runtime/debug"

	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/api/constants"
	"phl-skate-sharpening-api/core/models"
	"phl-skate-sharpening-api/utils"
)

type SharpeningCodeController struct {
	Log *logrus.Logger
}

type SharpeningCodeResponse struct {
	SharpeningCodes []SharpeningCodeInfoStruct `json:"SharpeningCode"`
}

type SharpeningCodeInfoStruct struct {
	ID          int             `json:"id"`
	Code        string          `json:"code"`
	StoreId     int             `json:"store_id"`
	IsValidCode bool            `json:"is_valid_code"`
	StoreInfo   StoreInfoStruct `json:"store_info"`
}

func (controller SharpeningCodeController) CheckIfCodeIsValid(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_PARAMS, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::SharpeningCodeController::CheckIfCodeIsValid - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	code := context.Params.ByName("code")

	codeInfo, isValid, err := context.Core.SharpeningCodeService.GetSharpeningCodeInfo(code)
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructSharpeningCodeResponse(codeInfo, isValid)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func ConstructSharpeningCodeResponse(codes []models.SharpeningCodeStruct, isValid bool) SharpeningCodeResponse {
	var response SharpeningCodeResponse
	sharpeningsCodeArray := make([]SharpeningCodeInfoStruct, 0)
	for _, codeInfo := range codes {
		codeResponse := ConstructSharpeningCodeStructResponse(codeInfo, isValid)
		sharpeningsCodeArray = append(sharpeningsCodeArray, codeResponse)
	}

	response.SharpeningCodes = sharpeningsCodeArray
	return response
}

func ConstructSharpeningCodeStructResponse(codeInfo models.SharpeningCodeStruct, isValid bool) SharpeningCodeInfoStruct {
	return SharpeningCodeInfoStruct{
		ID:          codeInfo.ID,
		Code:        codeInfo.Code,
		StoreId:     codeInfo.StoreId,
		IsValidCode: isValid,

		StoreInfo: ConstructStoreStructResponse(codeInfo.StoreInfo),
	}
}
