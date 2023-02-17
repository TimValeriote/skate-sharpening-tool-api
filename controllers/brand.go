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

type BrandController struct {
	Log *logrus.Logger
}

type BrandResponse struct {
	Brands []BrandInfoStruct `json:"Brands"`
}

type BrandInfoStruct struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	IsSkate   bool   `json:"is_skate"`
	IsSteel   bool   `json:"is_steel"`
	IsHolder  bool   `json:"is_holder"`
}

func (controller BrandController) GetBrands(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::BrandController::GetBrands - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	brands, err := context.Core.BrandService.GetAllBrands()
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructBrandResponse(brands)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func (controller BrandController) GetBrandById(writer http.ResponseWriter, request *http.Request) {
	context, err := utils.NewServiceFromContext(request, constants.CONTEXT_PARAMS, constants.CONTEXT_LOGGER, constants.CONTEXT_CORE)
	if err != nil {
		context.Log.WithFields(logrus.Fields{
			"event":      "phlapi::BrandController::GetBrandById - Failed to get value from context",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	brandId, err := strconv.Atoi(context.Params.ByName("brandId"))
	if err != nil {
		http.Error(writer, `{"error": "No valid id was provided"}`, http.StatusBadRequest)
		return
	}

	brand, err := context.Core.BrandService.GetBrandById(brandId)
	if err != nil {
		http.Error(writer, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response := ConstructBrandResponse(brand)

	context.Core.Commit()

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func ConstructBrandResponse(brands []models.BrandStruct) BrandResponse {
	var response BrandResponse
	brandsArray := make([]BrandInfoStruct, 0)
	for _, brand := range brands {
		brandResponse := ConstructBrandStructResponse(brand)
		brandsArray = append(brandsArray, brandResponse)
	}

	response.Brands = brandsArray
	return response
}

func ConstructBrandStructResponse(brand models.BrandStruct) BrandInfoStruct {
	return BrandInfoStruct{
		ID:        brand.ID,
		Name:      brand.Name,
		ShortName: brand.ShortName,
		IsSkate:   brand.IsSkate,
		IsSteel:   brand.IsSteel,
		IsHolder:  brand.IsHolder,
	}
}
