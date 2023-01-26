package core

import (
	"database/sql"
	//"fmt"
	"runtime/debug"
	//"strings"

	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/core/models"
)

type colourStore struct {
	database *models.CoreDatabase
	log      *logrus.Entry
}

func ColourStoreSetup(db *models.CoreDatabase, log *logrus.Entry) *colourStore {
	return &colourStore{
		database: db,
		log:      log,
	}
}

func (store *colourStore) GetAllColours() ([]models.ColourStruct, error) {
	sql := `SELECT id, colour FROM colour`
	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "colourStore::GetAllColours - Failed to prepare GetAllColours SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	ret, err := getColourFromQuery(query)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (store *colourStore) GetColourById(colourId int) ([]models.ColourStruct, error) {
	sql := `SELECT id, colour FROM colour WHERE id = ?`

	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "colourStore::GetColourById - Failed to prepare GetColourById SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query(colourId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	colours := make([]models.ColourStruct, 0)
	for rows.Next() {
		var colour models.ColourStruct
		err = rows.Scan(
			&colour.ID,
			&colour.Colour,
		)
		if err != nil {
			return nil, err
		}

		colours = append(colours, colour)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return colours, nil
}

func getColourFromQuery(query *sql.Stmt) ([]models.ColourStruct, error) {
	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	colours := make([]models.ColourStruct, 0)
	for rows.Next() {
		var colour models.ColourStruct
		err = rows.Scan(
			&colour.ID,
			&colour.Colour,
		)
		if err != nil {
			return nil, err
		}

		colours = append(colours, colour)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return colours, nil
}
