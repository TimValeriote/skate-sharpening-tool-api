package models

type ModelStruct struct {
	ID    int
	Name  string
	Alias string
}

type ModelService interface {
	GetAllModels() ([]ModelStruct, error)
	GetModelById(modelId int) ([]ModelStruct, error)
}
