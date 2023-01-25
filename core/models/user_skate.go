package models

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

type UserSkateService interface {
	GetAllUserSkatesByUserID(userId int) ([]UserSkateStruct, error)
	GetUserSkateByUserIdAndUserSkateId(userId int, userSkateId int) ([]UserSkateStruct, error)
}
