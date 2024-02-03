package models

type UserBank struct {
	BaseModel
	UserID      int    `json:"user_id"`
	Currency    string `json:"currency"`
	BankName    string `json:"bank_name"`
	Address     string `json:"address"`
	PostCode    string `json:"post_code"`
	AccountType string `json:"account_type"`
	Name        string `json:"name"`
	CardNumber  string `json:"card_number"`
	SwiftCode   string `json:"swift_code"`
	Branch      string `json:"branch"`
	BankCode    string `json:"bank_code"`
	Status      int    `json:"status"`
}

type UserBankDetail struct {
	UserBank
	AgentList []Agent `json:"agent_list"`
}
