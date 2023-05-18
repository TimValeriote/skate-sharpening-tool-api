package core

import (
	"runtime/debug"

	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/core/models"
)

type sharpeningCodeStore struct {
	database *models.CoreDatabase
	log      *logrus.Entry
}

func SharpeningCodeStoreSetup(db *models.CoreDatabase, log *logrus.Entry) *sharpeningCodeStore {
	return &sharpeningCodeStore{
		database: db,
		log:      log,
	}
}

func (store *sharpeningCodeStore) GetSharpeningCodeInfo(code string) ([]models.SharpeningCodeStruct, bool, error) {
	sql := `SELECT 
				cd.id, 
				cd.code, 
				cd.store_id,
				str.id, 
				str.name, 
				str.address, 
				str.city,  
				str.country,  
				str.phone_number, 
				str.store_number 
			FROM daily_sharpening_code AS cd 
			INNER JOIN store AS str ON cd.store_id = str.id
			WHERE code =  ?`

	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "sharpeningCodeStore::GetSharpeningCodeInfo - Failed to prepare GetSharpeningCodeInfo SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, false, err
	}
	defer query.Close()

	rows, err := query.Query(code)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()

	codes := make([]models.SharpeningCodeStruct, 0)
	for rows.Next() {
		var sharpeningCode models.SharpeningCodeStruct
		err = rows.Scan(
			&sharpeningCode.ID,
			&sharpeningCode.Code,
			&sharpeningCode.StoreId,

			&sharpeningCode.StoreInfo.ID,
			&sharpeningCode.StoreInfo.Name,
			&sharpeningCode.StoreInfo.Address,
			&sharpeningCode.StoreInfo.City,
			&sharpeningCode.StoreInfo.Country,
			&sharpeningCode.StoreInfo.PhoneNumber,
			&sharpeningCode.StoreInfo.StoreNumber,
		)
		if err != nil {
			return nil, false, err
		}

		codes = append(codes, sharpeningCode)
	}

	if rows.Err() != nil {
		return nil, false, err
	}

	validCode := false

	if len(codes) > 0 && len(codes) < 2 {
		validCode = true
	}

	return codes, validCode, nil
}

func (store *sharpeningCodeStore) InsertStoreCode(storeId int, code string) (err error) {
	sql := `INSERT INTO daily_sharpening_code (store_id, code) VALUES (?,?)`

	sqlStmt, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "sharpeningCodeStore::InsertStoreCode - Failed to prepare InsertStoreCode SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	_, err = sqlStmt.Exec(
		storeId,
		code,
	)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "sharpeningCodeStore::InsertStoreCode - Failed to execute InsertStoreCode SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	return
}
