package models

/*
    Status 0 = Reject
	Status 1 = Approval
	Status 2 = Padding Review
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

type WithdrawDetail struct {
	Withdraw
	CurrencyInfo Bank            `json:"info"`
	AgentList    []Agent         `json:"agent_list"`
	UserInfo     UserInformation `json:"user"`
}
