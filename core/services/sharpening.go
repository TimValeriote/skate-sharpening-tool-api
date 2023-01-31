package core

import (
	"database/sql"
	"fmt"
	"runtime/debug"
	//"strings"

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

	fmt.Println("yes sirrrrr")

	ret, err := getSharpeningsFromQuery(query, userId)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

/*func (store *sharpeningStore) GetUserById(userId int) (models.UsersStruct, error) {
	sql := `SELECT id, first_name, last_name, email, phone_number, uuid, is_staff FROM users WHERE id = ?`

	var userById models.UsersStruct

	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "sharpeningStore::GetUserById - Failed to prepare GetUserById SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return userById, err
	}
	defer query.Close()

	rows, err := query.Query(userId)
	if err != nil {
		return userById, err
	}
	defer rows.Close()

	users := make([]models.UsersStruct, 0)
	for rows.Next() {
		err = rows.Scan(
			&userById.ID,
			&userById.FirstName,
			&userById.LastName,
			&userById.Email,
			&userById.PhoneNumber,
			&userById.UUID,
			&userById.IsStaff,
		)
		if err != nil {
			return userById, err
		}

		users = append(users, userById)
	}

	if rows.Err() != nil {
		return userById, err
	}

	return userById, nil
}

func (store *sharpeningStore) CreateUser(userResponse *models.UsersStruct) (userId int, err error) {
	sql := `INSERT INTO users (
		first_name,
		last_name,
		email,
		phone_number,
		uuid
	) VALUES (?,?,?,?,?)`

	sqlStmt, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "sharpeningStore::CreateUser - Failed to prepare CreateUser SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	res, err := sqlStmt.Exec(
		userResponse.FirstName,
		userResponse.LastName,
		userResponse.Email,
		userResponse.PhoneNumber,
		userResponse.UUID,
	)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "sharpeningStore::CreateUser - Failed to execute CreateUser SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	insertedId, err := res.LastInsertId()
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "sharpeningStore::CreateUser - Failed to get the last inserted id",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	userId = int(insertedId)

	return
}*/

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
