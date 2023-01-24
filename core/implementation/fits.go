package implementation

import (
	"phl-skate-sharpening-api/core/models"
	services "phl-skate-sharpening-api/core/services"
)

type FitStore struct {
	*models.Core
}

func (store *FitStore) GetAllFits() ([]models.FitStruct, error) {
	return services.FitStoreSetup(store.Database, store.Log).GetAllFits()
}

func (store *FitStore) GetFitById(fitId int) ([]models.FitStruct, error) {
	return services.FitStoreSetup(store.Database, store.Log).GetFitById(fitId)
}
