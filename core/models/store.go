package models

type StoreStruct struct {
	ID          int
	Name        string
	Address     string
	City        string
	Country     string
	PhoneNumber string
	StoreNumber string
}

type StoreService interface {
	GetAllStores() ([]StoreStruct, error)
}
