package core

import (
	"database/sql"
	//"fmt"
	"runtime/debug"
	//"strings"

	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/core/models"
)

type fitStore struct {
	database *models.CoreDatabase
	log      *logrus.Entry
}

func FitStoreSetup(db *models.CoreDatabase, log *logrus.Entry) *fitStore {
	return &fitStore{
		database: db,
		log:      log,
	}
}

func (store *fitStore) GetAllFits() ([]models.FitStruct, error) {
	sql := `SELECT id, name FROM fits`
	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "fitStore::GetAllFits - Failed to prepare GetAllFits SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	ret, err := getFitsFromQuery(query)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (store *fitStore) GetFitById(fitId int) ([]models.FitStruct, error) {
	sql := `SELECT id, name FROM fits WHERE id = ?`

	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "fitStore::GetFitById - Failed to prepare GetFitById SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query(fitId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fits := make([]models.FitStruct, 0)
	for rows.Next() {
		var fit models.FitStruct
		err = rows.Scan(
			&fit.ID,
			&fit.Name,
		)
		if err != nil {
			return nil, err
		}

		fits = append(fits, fit)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return fits, nil
}

func getFitsFromQuery(query *sql.Stmt) ([]models.FitStruct, error) {
	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fits := make([]models.FitStruct, 0)
	for rows.Next() {
		var fit models.FitStruct
		err = rows.Scan(
			&fit.ID,
			&fit.Name,
		)
		if err != nil {
			return nil, err
		}

		fits = append(fits, fit)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return fits, nil
}
