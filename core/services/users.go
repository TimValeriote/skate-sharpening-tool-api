package core

import (
	"database/sql"
	//"fmt"
	"runtime/debug"
	//"strings"

	"github.com/sirupsen/logrus"
	"phl-skate-sharpening-api/core/models"
)

type userStore struct {
	database *models.CoreDatabase
	log      *logrus.Entry
}

func UserStoreSetup(db *models.CoreDatabase, log *logrus.Entry) *userStore {
	return &userStore{
		database: db,
		log:      log,
	}
}

func (store *userStore) GetAllUsers() ([]models.UsersStruct, error) {
	sql := `SELECT id, first_name, last_name, email, phone_number, uuid, is_staff FROM users`
	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userStore::GetAllUsers - Failed to prepare GetAllUsers SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	ret, err := getUsersFromQuery(query)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (store *userStore) GetUserById(userId int) ([]models.UsersStruct, error) {
	sql := `SELECT id, first_name, last_name, email, phone_number, uuid, is_staff FROM users WHERE id = ?`

	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userStore::GetUserById - Failed to prepare GetUserById SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]models.UsersStruct, 0)
	for rows.Next() {
		var user models.UsersStruct
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.PhoneNumber,
			&user.UUID,
			&user.IsStaff,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return users, nil
}

func getUsersFromQuery(query *sql.Stmt) ([]models.UsersStruct, error) {
	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]models.UsersStruct, 0)
	for rows.Next() {
		var user models.UsersStruct
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.PhoneNumber,
			&user.UUID,
			&user.IsStaff,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return users, nil
}
