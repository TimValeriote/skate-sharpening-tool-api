package models

type SharpeningCodeStruct struct {
	ID        int
	Code      string
	StoreId   int
	StoreInfo StoreStruct
}

type SharpeningCodeService interface {
	GetSharpeningCodeInfo(code string) ([]SharpeningCodeStruct, bool, error)
}
