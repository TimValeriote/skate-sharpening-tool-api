package implementation

import (
	"phl-skate-sharpening-api/core/models"
	services "phl-skate-sharpening-api/core/services"
)

type PeopleStore struct {
	*models.Core
}

func (store *PeopleStore) GetAllPeople() ([]models.PeopleStruct, error) {
	return services.PeopleStoreSetup(store.Database, store.Log).GetAllPeople()
}

func (store *PeopleStore) CreatePerson(personResponse *models.PeopleStruct) error {
	return services.PeopleStoreSetup(store.Database, store.Log).CreatePerson(personResponse)
}
