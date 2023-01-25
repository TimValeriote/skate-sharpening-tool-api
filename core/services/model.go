package core

import (
	"database/sql"
	//"fmt"
	"runtime/debug"
	//"strings"

	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/core/models"
)

type modelStore struct {
	database *models.CoreDatabase
	log      *logrus.Entry
}

func ModelStoreSetup(db *models.CoreDatabase, log *logrus.Entry) *modelStore {
	return &modelStore{
		database: db,
		log:      log,
	}
}

func (store *modelStore) GetAllModels() ([]models.ModelStruct, error) {
	sql := `SELECT 
					model.id, 
					model.name, 
					model.alias,
			FROM model`
	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "modelStore::GetAllModels - Failed to prepare GetAllModels SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	ret, err := getModelsFromQuery(query)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (store *modelStore) GetModelById(modelId int) ([]models.ModelStruct, error) {
	sql := `SELECT 
					model.id, 
					model.name, 
					model.alias,
			FROM model
			WHERE model.id = ?`

	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "modelStore::GetModelById - Failed to prepare GetModelById SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query(modelId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	modelsArray := make([]models.ModelStruct, 0)
	for rows.Next() {
		var model models.ModelStruct
		err = rows.Scan(
			&model.ID,
			&model.Name,
			&model.Alias,
		)
		if err != nil {
			return nil, err
		}
		modelsArray = append(modelsArray, model)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return modelsArray, nil
}

func getModelsFromQuery(query *sql.Stmt) ([]models.ModelStruct, error) {
	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	modelsArray := make([]models.ModelStruct, 0)
	for rows.Next() {
		var model models.ModelStruct
		err = rows.Scan(
			&model.ID,
			&model.Name,
			&model.Alias,
		)
		if err != nil {
			return nil, err
		}
		modelsArray = append(modelsArray, model)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return modelsArray, nil
}
