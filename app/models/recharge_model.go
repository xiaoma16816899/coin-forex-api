package models

// 充值
type Recharge struct {
	BaseModel
	UserID     int    `json:"user_id"`
	Virtual    int    `json:"virtual"`
	Mode       int    `json:"mode"`
	Title      string `json:"title"`
	OrderNo    string `json:"order_no"`
	Currency   string `json:"currency"`
	Rate       string `json:"rate"`
	Amount     string `json:"amount"`
	RealAmount string `json:"real_amount"`
	Img        string `json:"img"`
	Reason     string `json:"reason"`
	Remark     string `json:"remark"`
	Times      int    `json:"times"`
	Operator   string `json:"operator"`
	ReviewTime int    `json:"review_time"`
	Status     int    `json:"status"`
}

type RechargeDetail struct {
	Recharge
	User      UserInfo     `json:"user"`
	Info      InfoCurrency `json:"info"`
	AgentList []Agent      `json:"agent_list"`
}
