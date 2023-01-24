package models

type StoreStruct struct {
	ID          int
	Name        string
	Address     string
	City        string
	Country     string
	PhoneNumber string
}

type StoreService interface {
	GetAllStores() ([]StoreStruct, error)
}
