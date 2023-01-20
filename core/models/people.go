package models

type PeopleStruct struct {
	FirstName string
	LastName  string
}

type PeopleService interface {
	GetAllPeople() ([]PeopleStruct, error)
	CreatePerson(personResponse *PeopleStruct) error
}
