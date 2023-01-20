package models

type PeopleStruct struct {
	ID        int
	FirstName string
	LastName  string
}

type PeopleService interface {
	GetAllPeople() ([]PeopleStruct, error)
	GetPersonByID(personID int) ([]PeopleStruct, error)
	CreatePerson(personResponse *PeopleStruct) error
	UpdatePerson(person *PeopleStruct) error
}
