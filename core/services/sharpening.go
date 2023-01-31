package core

import (
	"database/sql"
	"runtime/debug"

	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/core/models"
)

type sharpeningStore struct {
	database *models.CoreDatabase
	log      *logrus.Entry
}

func SharpeningStoreSetup(db *models.CoreDatabase, log *logrus.Entry) *sharpeningStore {
	return &sharpeningStore{
		database: db,
		log:      log,
	}
}

func (store *sharpeningStore) GetOpenSharpeningsForUser(userId int) ([]models.SharpeningStruct, error) {
	sql := `SELECT id, user_id, user_skate_id, store_id FROM open_sharpenings WHERE user_id = ?`
	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "sharpeningStore::GetOpenSharpeningsForUser - Failed to prepare GetOpenSharpeningsForUser SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	ret, err := getSharpeningsFromQuery(query, userId)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func getSharpeningsFromQuery(query *sql.Stmt, userId int) ([]models.SharpeningStruct, error) {
	rows, err := query.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sharpenings := make([]models.SharpeningStruct, 0)
	for rows.Next() {
		var sharpening models.SharpeningStruct
		err = rows.Scan(
			&sharpening.ID,
			&sharpening.UserId,
			&sharpening.UserSkateId,
			&sharpening.StoreId,
		)
		if err != nil {
			return nil, err
		}

		sharpenings = append(sharpenings, sharpening)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return sharpenings, nil
}
