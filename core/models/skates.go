package models

type SkateStruct struct {
	ID    int
	Model ModelStruct
	Brand BrandStruct
}

type SkateService interface {
	GetAllSkates() ([]SkateStruct, error)
}
