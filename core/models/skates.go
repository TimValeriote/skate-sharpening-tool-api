package models

type SkateStruct struct {
	ID    int
	Model ModelStruct
	Brand BrandStruct
	Fit   FitStruct
}

type SkateService interface {
	GetAllSkates() ([]SkateStruct, error)
}
