package implementation

import (
	"phl-skate-sharpening-api/core/models"
	services "phl-skate-sharpening-api/core/services"
)

type UserSkateStore struct {
	*models.Core
}

func (store *UserSkateStore) GetAllUserSkatesByUserID(userID int) ([]models.UserSkateStruct, error) {
	return services.UserSkateStoreSetup(store.Database, store.Log).GetAllUserSkatesByUserID(userID)
}
