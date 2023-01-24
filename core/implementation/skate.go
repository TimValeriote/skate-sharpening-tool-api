package implementation

import (
	"phl-skate-sharpening-api/core/models"
	services "phl-skate-sharpening-api/core/services"
)

type SkateStore struct {
	*models.Core
}

func (store *SkateStore) GetAllSkates() ([]models.SkateStruct, error) {
	return services.SkateStoreSetup(store.Database, store.Log).GetAllSkates()
}
