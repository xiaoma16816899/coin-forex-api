package models

/*
	Mode: 线下, 彩金
	Mode: 0 ( Offline )
	Mode: 1 ( Bonus )
	Status 0 = Reject for deposit
	Status 1 = Pass for deposit
	Status 2 = Padding Review
*/
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
	Times      int    `json:"times"` // times of recharge
	Operator   string `json:"operator"`
	ReviewTime int    `json:"review_time"`
	Status     int    `json:"status"`
}

type RechargeDetail struct {
	Recharge
	UserInfo  UserInformation `json:"user"`
	Info      InfoCurrency    `json:"info"`
	AgentList []Agent         `json:"agent_list"`
}
