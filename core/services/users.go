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

func (store *userStore) GetUserById(userId int) (models.UsersStruct, error) {
	sql := `SELECT id, first_name, last_name, email, phone_number, uuid, is_staff FROM users WHERE id = ?`

	var userById models.UsersStruct

	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userStore::GetUserById - Failed to prepare GetUserById SELECT query.",
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

func (store *userStore) GetUserByEmail(userEmail string) (models.UsersStruct, error) {
	sql := `SELECT id, first_name, last_name, email, phone_number, uuid, is_staff FROM users WHERE email = ?`

	var userByEmail models.UsersStruct

	query, err := store.database.Tx.Prepare(sql)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userStore::GetUserByEmail - Failed to prepare GetUserByEmail SELECT query.",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return userByEmail, err
	}
	defer query.Close()

	rows, err := query.Query(userEmail)
	if err != nil {
		return userByEmail, err
	}
	defer rows.Close()

	users := make([]models.UsersStruct, 0)
	for rows.Next() {
		err = rows.Scan(
			&userByEmail.ID,
			&userByEmail.FirstName,
			&userByEmail.LastName,
			&userByEmail.Email,
			&userByEmail.PhoneNumber,
			&userByEmail.UUID,
			&userByEmail.IsStaff,
		)
		if err != nil {
			return userByEmail, err
		}

		users = append(users, userByEmail)
	}

	if rows.Err() != nil {
		return userByEmail, err
	}

	return userByEmail, nil
}

func (store *userStore) CreateUser(userResponse *models.UsersStruct) (userId int, err error) {
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
			"event":      "userStore::CreateUser - Failed to prepare CreateUser SQL",
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
			"event":      "userStore::CreateUser - Failed to execute CreateUser SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	insertedId, err := res.LastInsertId()
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userStore::CreateUser - Failed to get the last inserted id",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	userId = int(insertedId)

	return
}

func (store *userStore) UpdateUser(user models.UpdateUserStruct) (result sql.Result, err error) {
	sql := `UPDATE users SET 
					first_name = ?, 
					last_name = ?,
					email = ?,
					phone_number = ?
			WHERE id = ?`

	result, err = store.database.Tx.Exec(sql,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PhoneNumber,
		user.UserID)
	if err != nil {
		store.log.WithFields(logrus.Fields{
			"event":      "userStore::UpdateUser - Failed to execute UpdateUser SQL",
			"query":      sql,
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return
	}

	return
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
