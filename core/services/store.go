package core

import (
	"database/sql"
	"runtime/debug"

	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/core/models"
)

type storeStore struct {
	database *models.CoreDatabase
	log      *logrus.Entry
}

func StoreStoreSetup(db *models.CoreDatabase, log *logrus.Entry) *storeStore {
	return &storeStore{
		database: db,
		log:      log,
	}
}

func (store *storeStore) GetAllStores() ([]models.StoreStruct, error) {
	sql := `SELECT id, name, address, city, country, phone_number, store_number FROM store`
	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "storeStore::GetAllStores - Failed to prepare GetAllStores SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	ret, err := getStoreFromQuery(query)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (store *storeStore) GetStoreById(storeId int) ([]models.StoreStruct, error) {
	sql := `SELECT id, name, address, city, country, phone_number, store_number FROM store WHERE id = ?`

	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "storeStore::GetStoreById - Failed to prepare GetStoreById SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query(storeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stores := make([]models.StoreStruct, 0)
	for rows.Next() {
		var store models.StoreStruct
		err = rows.Scan(
			&store.ID,
			&store.Name,
			&store.Address,
			&store.City,
			&store.Country,
			&store.PhoneNumber,
			&store.StoreNumber,
		)
		if err != nil {
			return nil, err
		}

		stores = append(stores, store)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return stores, nil
}

func getStoreFromQuery(query *sql.Stmt) ([]models.StoreStruct, error) {
	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stores := make([]models.StoreStruct, 0)
	for rows.Next() {
		var store models.StoreStruct
		err = rows.Scan(
			&store.ID,
			&store.Name,
			&store.Address,
			&store.City,
			&store.Country,
			&store.PhoneNumber,
			&store.StoreNumber,
		)
		if err != nil {
			return nil, err
		}

		stores = append(stores, store)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return stores, nil
}
