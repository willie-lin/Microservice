package models

type Company struct {
	Ticker string `json:"ticker" form:"ticker" query:"ticker"`
}
