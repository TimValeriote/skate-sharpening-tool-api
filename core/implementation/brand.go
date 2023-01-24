package implementation

import (
	"phl-skate-sharpening-api/core/models"
	services "phl-skate-sharpening-api/core/services"
)

type BrandStore struct {
	*models.Core
}

func (store *BrandStore) GetAllBrands() ([]models.BrandStruct, error) {
	return services.BrandStoreSetup(store.Database, store.Log).GetAllBrands()
}

func (store *BrandStore) GetBrandById(brandId int) ([]models.BrandStruct, error) {
	return services.BrandStoreSetup(store.Database, store.Log).GetBrandById(brandId)
}
