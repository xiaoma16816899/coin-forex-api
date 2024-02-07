package models

/*
Status 0 = Reject Load
Status 1 = Done For Load Payment
Status 2 = Padding Review
Status 3 = Approve Load
*/
type Borrow struct {
	BaseModel
	UserID     int    `json:"user_id"`
	Title      string `json:"title"`
	Day        int    `json:"day"`
	Amount     string `json:"amount"`
	Rate       string `json:"rate"`
	FwImg      string `json:"fw_img"`
	SrImg      string `json:"sr_img"`
	YhImg      string `json:"yh_img"`
	ZjImg      string `json:"zj_img"`
	Reason     string `json:"reason"`
	ReviewTime int    `json:"review_time"`
	BackTime   int    `json:"back_time"`
	Operator   string `json:"operator"` // Represent the admin username that approval or reject
	Status     int    `json:"status"`
}

type BorrowDetail struct {
	Borrow
	UserInfo  UserInformation `json:"user"`
	AgentList []Agent         `json:"agent_list"`
}
