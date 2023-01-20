package models

import (
	"context"
	"database/sql"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

func (model *Core) SetLogger(log *logrus.Logger) {
	model.SetLoggerEntry(log.WithFields(logrus.Fields{}))
}

func (model *Core) SetLoggerEntry(log *logrus.Entry) {
	model.Log = log
}

func (model *Core) Begin() error {
	var err error

	model.Log.Info("Starting Core Transaction")
	model.Database.Tx, err = model.Database.DB.BeginTx(model.Ctx, nil)
	if err != nil {
		model.Log.WithFields(logrus.Fields{
			"event":      "Core::Begin - Failed to Begin transaction",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
	}
	return err
}

func (model *Core) Commit() error {
	var err error

	if model.Database.Tx != nil {
		model.Log.Info("Committing Core Transaction")
		err = model.Database.Tx.Commit()
		if err != nil && err != sql.ErrTxDone {
			model.Log.WithFields(logrus.Fields{
				"event":      "Core::Commit - Failed to Commit transaction",
				"stackTrace": string(debug.Stack()),
			}).Error(err)
			return err
		}
	}

	return nil
}

func (model *Core) Rollback() error {
	var err error

	if model.Database.Tx != nil {
		model.Log.Info("Rolling Back Core Transaction")
		err = model.Database.Tx.Rollback()
		if err != nil && err != sql.ErrTxDone {
			model.Log.WithFields(logrus.Fields{
				"event":      "Core::Rollback - Failed to Rollback transaction",
				"stackTrace": string(debug.Stack()),
			}).Error(err)
			return err
		}
	}

	return nil
}

func (model *Core) Close() error {
	var err error
	if model.Database != nil {
		err = model.Rollback()
		if err == nil || err == sql.ErrTxDone {
			if model.Database.DB != nil {
				err = model.Database.DB.Close()
			}
		}
	}

	if err != nil && err != sql.ErrTxDone {
		model.Log.WithFields(logrus.Fields{
			"event":      "Core::Close - Failed to Close database connection",
			"stackTrace": string(debug.Stack()),
		}).Error(err)
		return err
	}

	return nil
}

type WhereQueryStruct struct {
	Generate     func(map[string][]interface{}) (string, []interface{})
	Placeholders map[string][]interface{}
}

type CoreDatabase struct {
	DB *sql.DB
	Tx *sql.Tx
}

type Core struct {
	Database *CoreDatabase
	Log      *logrus.Entry
	Ctx      context.Context

	PeopleService PeopleService
}
