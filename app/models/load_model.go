package models

type Load struct {
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
	Operator   string `json:"operator"`
	Status     int    `json:"status"`
}

type LoadDetail struct {
	Load
	User      UserInfo `json:"user"`
	AgentList []Agent  `json:"agent_list"`
}
