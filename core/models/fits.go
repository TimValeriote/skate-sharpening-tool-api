package models

type FitStruct struct {
	ID   int
	Name string
}

type FitService interface {
	GetAllFits() ([]FitStruct, error)
	GetFitById(fitID int) ([]FitStruct, error)
}
