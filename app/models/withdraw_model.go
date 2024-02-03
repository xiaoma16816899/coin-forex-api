package models

/*
	Status 
*/
// 提现列表
type Withdraw struct {
	BaseModel
	UserID      int    `json:"user_id"`
	Virtual     int    `json:"virtual"`
	Mode        int    `json:"mode"`
	Title       string `json:"title"`
	OrderNo     string `json:"order_no"`
	Currency    string `json:"currency"`
	Rate        string `json:"rate"`
	Amount      string `json:"amount"`
	ApplyAmount string `json:"apply_amount"`
	RealAmount  string `json:"real_amount"`
	Fee         string `json:"fee"`
	Reason      string `json:"reason"`
	Remark      string `json:"remark"`
	Operator    string `json:"operator"`
	ReviewTime  int    `json:"review_time"`
	CreateIP    string `json:"create_ip"`
	Status      int    `json:"status"`
	Nums        int    `json:"nums"`
}

type Info struct {
	Fn          string `json:"fn"`
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

type WithdrawDetail struct {
	Withdraw
	CurrencyInfo Info     `json:"info"`
	AgentList    []Agent  `json:"agent_list"`
	User         UserInfo `json:"user"`
}
