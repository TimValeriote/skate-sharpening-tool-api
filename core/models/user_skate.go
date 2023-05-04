package models

import (
	"database/sql"
)

type UserSkateStruct struct {
	ID              int
	Skate           SkateStruct
	Holder          BrandStruct
	HolderSize      float32
	SkateSize       float32
	LaceColour      ColourStruct
	HasSteel        bool
	Steel           BrandStruct
	HasGuards       bool
	GuardColour     ColourStruct
	PreferredRadius string
	Fit             FitStruct
}

type CreateUserSkateStruct struct {
	UserID          int
	SkateID         string
	HolderBrandID   string
	HolderSize      string
	SkateSize       string
	LaceColourID    string
	HasSteel        string
	SteelID         string
	HasGuards       string
	GuardColourID   string
	PreferredRadius string
	FitID           string
}

type UpdateUserSkateStruct struct {
	UserSkateID     int
	SkateID         string
	HolderBrandID   string
	HolderSize      string
	SkateSize       string
	LaceColourID    string
	HasSteel        string
	SteelID         string
	HasGuards       string
	GuardColourID   string
	PreferredRadius string
	FitID           string
}

type UserSkateService interface {
	GetAllUserSkatesByUserID(userId int) ([]UserSkateStruct, error)
	GetUserSkateByUserIdAndUserSkateId(userId int, userSkateId int) (UserSkateStruct, error)
	CreateUserSkate(userSkate *CreateUserSkateStruct) (int, error)
	UpdateUserSkate(userSkate UpdateUserSkateStruct) (sql.Result, error)
	DeleteUserSkate(userId int, userSkateId int) (sql.Result, error)
	GetUserSkatesNotBeingSharpened(userId int) ([]UserSkateStruct, error)
}
