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
	GetUserByEmail(userEmail string) ([]UsersStruct, error)
}
