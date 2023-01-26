package implementation

import (
	"database/sql"

	"phl-skate-sharpening-api/core/models"
	services "phl-skate-sharpening-api/core/services"
)

type UserSkateStore struct {
	*models.Core
}

func (store *UserSkateStore) GetAllUserSkatesByUserID(userID int) ([]models.UserSkateStruct, error) {
	return services.UserSkateStoreSetup(store.Database, store.Log).GetAllUserSkatesByUserID(userID)
}

func (store *UserSkateStore) GetUserSkateByUserIdAndUserSkateId(userID int, userSkateId int) ([]models.UserSkateStruct, error) {
	return services.UserSkateStoreSetup(store.Database, store.Log).GetUserSkateByUserIdAndUserSkateId(userID, userSkateId)
}

func (store *UserSkateStore) CreateUserSkate(userSkateResponse *models.CreateUserSkateStruct) (int, error) {
	return services.UserSkateStoreSetup(store.Database, store.Log).CreateUserSkate(userSkateResponse)
}

func (store *UserSkateStore) UpdateUserSkate(userSkate models.UpdateUserSkateStruct) (sql.Result, error) {
	return services.UserSkateStoreSetup(store.Database, store.Log).UpdateUserSkate(userSkate)
}
