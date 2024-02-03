package models

type InfoCurrency struct {
	Fn       string `json:"fn"`
	Currency string `json:"currency"`
	Symbol   string `json:"symbol"`
	Rate     int    `json:"rate"`
	Fixed    int    `json:"fixed"`
}
