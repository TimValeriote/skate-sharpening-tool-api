package models

import (
	"database/sql"
)

type UsersStruct struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	UUID        string
	IsStaff     bool
}

type UpdateUserStruct struct {
	UserID      int
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
}

type UserService interface {
	GetAllUsers() ([]UsersStruct, error)
	GetUserById(userId int) ([]UsersStruct, error)
	CreateUser(user *UsersStruct) (int, error)
	UpdateUser(user UpdateUserStruct) (sql.Result, error)
}
