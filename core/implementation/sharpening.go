package implementation

import (
	//"database/sql"

	"phl-skate-sharpening-api/core/models"
	services "phl-skate-sharpening-api/core/services"
)

type SharpeningStore struct {
	*models.Core
}

func (store *SharpeningStore) GetOpenSharpeningsForUser(userId int) ([]models.SharpeningStruct, error) {
	return services.SharpeningStoreSetup(store.Database, store.Log).GetOpenSharpeningsForUser(userId)
}

/*func (store *SharpeningStore) CreateSharpeningForUser(userResponse *models.UsersStruct) (int, error) {
	return services.SharpeningStoreSetup(store.Database, store.Log).CreateUser(userResponse)
}

func (store *SharpeningStore) RemoveSharpeningForUser(sharpeningId int) (sql.Result, error) {
	return services.SharpeningStoreSetup(store.Database, store.Log).UpdateUser(user)
}*/
