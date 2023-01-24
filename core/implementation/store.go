package implementation

import (
	"phl-skate-sharpening-api/core/models"
	services "phl-skate-sharpening-api/core/services"
)

type StoreStore struct {
	*models.Core
}

func (store *StoreStore) GetAllStores() ([]models.StoreStruct, error) {
	return services.StoreStoreSetup(store.Database, store.Log).GetAllStores()
}
