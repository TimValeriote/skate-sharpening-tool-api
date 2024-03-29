package implementation

import (
	"database/sql"

	"phl-skate-sharpening-api/core/models"
	services "phl-skate-sharpening-api/core/services"
)

type SharpeningStore struct {
	*models.Core
}

func (store *SharpeningStore) GetOpenSharpeningsForUser(userId int) ([]models.SharpeningStruct, error) {
	return services.SharpeningStoreSetup(store.Database, store.Log).GetOpenSharpeningsForUser(userId)
}

func (store *SharpeningStore) DeleteSharpen(sharpenId int, userId int) (sql.Result, error) {
	return services.SharpeningStoreSetup(store.Database, store.Log).DeleteSharpen(sharpenId, userId)
}

func (store *SharpeningStore) AddSharpening(userId int, userSkateId int, storeId int) (err error) {
	return services.SharpeningStoreSetup(store.Database, store.Log).AddSharpening(userId, userSkateId, storeId)
}
