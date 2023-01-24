package models

type BrandStruct struct {
	ID        int
	Name      string
	ShortName string
	IsSkate   bool
	IsSteel   bool
	IsHolder  bool
}

type BrandService interface {
	GetAllBrands() ([]BrandStruct, error)
	GetBrandById(brandId int) ([]BrandStruct, error)
}
