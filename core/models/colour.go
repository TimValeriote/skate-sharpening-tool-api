package models

type ColourStruct struct {
	ID     int
	Colour string
}

type ColourService interface {
	GetAllColours() ([]ColourStruct, error)
	GetColourById(colourId int) ([]ColourStruct, error)
}
