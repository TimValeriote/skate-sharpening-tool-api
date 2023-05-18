package implementation

import (
	"phl-skate-sharpening-api/core/models"
	services "phl-skate-sharpening-api/core/services"
)

type SharpeningCodeStore struct {
	*models.Core
}

func (store *SharpeningCodeStore) GetSharpeningCodeInfo(code string) ([]models.SharpeningCodeStruct, bool, error) {
	return services.SharpeningCodeStoreSetup(store.Database, store.Log).GetSharpeningCodeInfo(code)
}

func (store *SharpeningCodeStore) InsertStoreCode(storeId int, code string) (err error) {
	return services.SharpeningCodeStoreSetup(store.Database, store.Log).InsertStoreCode(storeId, code)
}
