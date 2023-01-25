package models

type UsersStruct struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	UUID        string
	IsStaff     bool
}

type UserService interface {
	GetAllUsers() ([]UsersStruct, error)
	GetUserById(userId int) ([]UsersStruct, error)
}
