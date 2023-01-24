package implementation

import (
	"phl-skate-sharpening-api/core/models"
	services "phl-skate-sharpening-api/core/services"
)

type UserStore struct {
	*models.Core
}

func (store *UserStore) GetAllUsers() ([]models.UsersStruct, error) {
	return services.UserStoreSetup(store.Database, store.Log).GetAllUsers()
}

func (store *UserStore) GetUserByEmail(userEmail string) ([]models.UsersStruct, error) {
	return services.UserStoreSetup(store.Database, store.Log).GetUserByEmail(userEmail)
}
