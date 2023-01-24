package models

type ModelStruct struct {
	ID    int
	Name  string
	Alias string
	Brand BrandStruct
}

type ModelService interface {
	GetAllModels() ([]ModelStruct, error)
	GetModelById(modelId int) ([]ModelStruct, error)
}
