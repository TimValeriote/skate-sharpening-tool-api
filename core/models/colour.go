package models

import (
	"gopkg.in/guregu/null.v3"
)

type ColourStruct struct {
	ID     null.Int
	Colour null.String
}

type ColourService interface {
	GetAllColours() ([]ColourStruct, error)
	GetColourByName(colourName string) ([]ColourStruct, error)
}
