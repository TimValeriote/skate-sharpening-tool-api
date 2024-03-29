package implementation

import (
	"phl-skate-sharpening-api/core/models"
	services "phl-skate-sharpening-api/core/services"
)

type ColourStore struct {
	*models.Core
}

func (store *ColourStore) GetAllColours() ([]models.ColourStruct, error) {
	return services.ColourStoreSetup(store.Database, store.Log).GetAllColours()
}

func (store *ColourStore) GetColourById(colourId int) ([]models.ColourStruct, error) {
	return services.ColourStoreSetup(store.Database, store.Log).GetColourById(colourId)
}
