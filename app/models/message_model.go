package models

// 站内信
type Message struct {
	UserID   int    `json:"user_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	IsRead   int    `json:"is_read"`
	ReadNums int    `json:"read_nums"`
	Operator string `json:"operator"`
	Status   int    `json:"status"`
}

type MessageDetail struct {
	Message
	UserInfo  UserInformation `json:"user"`
	AgentList []Agent         `json:"agent_list"`
}
