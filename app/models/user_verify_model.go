package models

// 实名认证
type UserVerify struct {
	BaseModel
	UserID     int    `json:"user_id"`
	Type       int    `json:"type"`
	Country    string `json:"country"`
	Name       string `json:"name"`
	IDCard     string `json:"id_card"`
	FrontImg   string `json:"front_img"`
	ReverseImg string `json:"reverse_img"`
	Reason     string `json:"reason"`
	Operator   string `json:"operator"`
	Status     int    `json:"status"`
}

type UserVerifyDetail struct {
	UserVerify
	UserInfo  UserInformation `json:"user"`
	AgentList []Agent         `json:"agent_list"`
}
