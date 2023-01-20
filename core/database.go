package core

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const driverName string = "mysql"
const MYSQL_KEY_EXISTS = 1062

func ConnectToDatabase(host, port, username, password, databaseName string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&collation=utf8mb4_unicode_ci", username, password, host, port, databaseName)

	db, err = sql.Open(driverName, connectString)
	if err != nil {
		return db, err
	}

	// Ensure that we can actually talk to the database. sql.Open() doesn't test this, need to actually ping.
	err = db.Ping()
	if err != nil {
		return db, err
	}
	db.SetMaxIdleConns(0)

	return db, nil
}

func CloseDatabase(db *sql.DB) error {
	var err error
	if db != nil {
		err = db.Close()
	}
	return err
}
