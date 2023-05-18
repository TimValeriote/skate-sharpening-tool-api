package models

import (
	"database/sql"
)

type SharpeningStruct struct {
	ID           int
	UserId       int
	UserSkateId  int
	StoreId      int
	ProgressText string
}

type SharpeningService interface {
	GetOpenSharpeningsForUser(userId int) ([]SharpeningStruct, error)
	DeleteSharpen(sharpenId int, userId int) (sql.Result, error)
	AddSharpening(userId int, userSkateId int, storeId int) (err error)
}
