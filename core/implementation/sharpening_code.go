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
