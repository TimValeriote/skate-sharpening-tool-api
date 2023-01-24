package implementation

import (
	"phl-skate-sharpening-api/core/models"
	services "phl-skate-sharpening-api/core/services"
)

type ModelStore struct {
	*models.Core
}

func (store *ModelStore) GetAllModels() ([]models.ModelStruct, error) {
	return services.ModelStoreSetup(store.Database, store.Log).GetAllModels()
}

func (store *ModelStore) GetModelById(modelId int) ([]models.ModelStruct, error) {
	return services.ModelStoreSetup(store.Database, store.Log).GetModelById(modelId)
}
