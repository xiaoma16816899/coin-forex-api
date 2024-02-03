package models

type UserVerify struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	Type       int    `json:"type"`
	Country    string `json:"country"`
	Name       string `json:"name"`
	IDCard     string `json:"id_card"`
	FrontImg   string `json:"front_img"`
	ReverseImg string `json:"reverse_img"`
	Reason     string `json:"reason"`
	CreateTime int    `json:"create_time"`
	ModifyTime int    `json:"modify_time"`
	Del        int    `json:"del"`
	Operator   string `json:"operator"`
	Status     int    `json:"status"`
}

type UserVerifyDetail struct {
	UserVerify
	User      UserInfo `json:"user"`
	AgentList []Agent  `json:"agent_list"`
}
