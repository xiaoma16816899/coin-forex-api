package models

type Bank struct {
	Fn          string `json:"fn"` // Represent the wallet, system
	BankName    string `json:"bank_name"`
	CardNumber  string `json:"card_number"`
	Name        string `json:"name"`
	BankCode    string `json:"bank_code"`
	Branch      string `json:"branch"`
	Currency    string `json:"currency"`
	SwiftCode   string `json:"swift_code"`
	Address     string `json:"address"`
	PostCode    string `json:"post_code"`
	AccountType string `json:"account_type"`
	Symbol      string `json:"symbol"`
	Rate        string `json:"rate"`
	Fee         string `json:"fee"`
	Fixed       string `json:"fixed"`
}
