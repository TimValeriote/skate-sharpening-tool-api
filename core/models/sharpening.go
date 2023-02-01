package models

import (
	"database/sql"
)

type SharpeningStruct struct {
	ID          int
	UserId      int
	UserSkateId int
	StoreId     int
}

type SharpeningService interface {
	GetOpenSharpeningsForUser(userId int) ([]SharpeningStruct, error)
	DeleteSharpen(sharpenId int, userId int) (sql.Result, error)
	//CreateSharpeningForUser(sharpening *SharpeningStruct) (int, error)
	//RemoveSharpeningForUser(sharpeningId int) (sql.Result, error)
}
