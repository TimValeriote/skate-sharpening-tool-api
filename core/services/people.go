package core

import (
	"database/sql"
	//"fmt"
	"runtime/debug"
	//"strings"

	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/core/models"
)

type peopleStore struct {
	database *models.CoreDatabase
	log      *logrus.Entry
}

func PeopleStoreSetup(db *models.CoreDatabase, log *logrus.Entry) *peopleStore {
	return &peopleStore{
		database: db,
		log:      log,
	}
}

func (store *peopleStore) GetAllPeople() ([]models.PeopleStruct, error) {
	sql := `SELECT id, first_name, last_name FROM people`

	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "peopleStore::GetAllPeople - Failed to prepare GetAllPeople SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	ret, err := getPeopleFromQuery(query)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (store *peopleStore) GetPersonByID(personID int) ([]models.PeopleStruct, error) {
	sql := `SELECT id, first_name, last_name FROM people WHERE id = ?`

	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "peopleStore::GetPersonByID - Failed to prepare GetPersonByID SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query(personID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	people := make([]models.PeopleStruct, 0)
	for rows.Next() {
		var person models.PeopleStruct
		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName)
		if err != nil {
			return nil, err
		}

		people = append(people, person)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return people, nil
}

func (store *peopleStore) CreatePerson(personResponse *models.PeopleStruct) (err error) {
	sql := `INSERT INTO people(first_name, last_name) VALUES (?,?)`

	sqlStmt, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "peopleStore::CreatePerson - Failed to prepare CreatePerson SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return err
	}

	_, err = sqlStmt.Exec(personResponse.FirstName, personResponse.LastName)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "peopleStore::CreatePerson - Failed to execute CreatePerson SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return err
	}

	return nil
}

func (store *peopleStore) UpdatePerson(person *models.PeopleStruct) (err error) {
	sql := `UPDATE people SET first_name = ?, last_name = ? WHERE id = ?`

	sqlStmt, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "peopleStore::UpdatePerson - Failed to prepare UpdatePerson SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return err
	}

	_, err = sqlStmt.Exec(person.FirstName, person.LastName, person.ID)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "peopleStore::UpdatePerson - Failed to execute UpdatePerson SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return err
	}

	return nil
}

func getPeopleFromQuery(query *sql.Stmt) ([]models.PeopleStruct, error) {
	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	people := make([]models.PeopleStruct, 0)
	for rows.Next() {
		var person models.PeopleStruct
		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName)
		if err != nil {
			return nil, err
		}

		people = append(people, person)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return people, nil
}
