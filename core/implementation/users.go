package implementation

import (
	"database/sql"

	"phl-skate-sharpening-api/core/models"
	services "phl-skate-sharpening-api/core/services"
)

type UserStore struct {
	*models.Core
}

func (store *UserStore) GetAllUsers() ([]models.UsersStruct, error) {
	return services.UserStoreSetup(store.Database, store.Log).GetAllUsers()
}

func (store *UserStore) GetUserById(userId int) ([]models.UsersStruct, error) {
	return services.UserStoreSetup(store.Database, store.Log).GetUserById(userId)
}

func (store *UserStore) CreateUser(userResponse *models.UsersStruct) (int, error) {
	return services.UserStoreSetup(store.Database, store.Log).CreateUser(userResponse)
}

func (store *UserStore) UpdateUser(user models.UpdateUserStruct) (sql.Result, error) {
	return services.UserStoreSetup(store.Database, store.Log).UpdateUser(user)
}
