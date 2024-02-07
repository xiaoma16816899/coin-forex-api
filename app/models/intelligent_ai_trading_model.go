package models

/*
	Status 0 = Rejects
	Status 1 = Padding Review
	Status 2 = Approval
*/
type IntelligentAITrading struct {
	BaseModel
	UserID       int    `json:"user_id"`
	MiningID     int    `json:"mining_id"`
	OrderNo      string `json:"order_no"`
	Title        string `json:"title"`
	Amount       string `json:"amount"`
	ProfitAmount string `json:"profit_amount"`
	Day          int    `json:"day"`
	NowRate      string `json:"now_rate"`
	MinRate      string `json:"min_rate"`
	MaxRate      string `json:"max_rate"`
	BcRate       string `json:"bc_rate"`
	BcAmount     string `json:"bc_amount"`
	EndTime      int    `json:"end_time"`
	BackTime     int    `json:"back_time"`
	TaskTime     int    `json:"task_time"`
	Reason       string `json:"reason"`
	Operator     string `json:"operator"`
	Status       int    `json:"status"`
	LastDay      int    `json:"lastDay"`
}

type IntelligentAITradingDetail struct {
	IntelligentAITrading
	UserInfo  UserInformation `json:"user"`
	AgentList []Agent         `json:"agent_list"`
}
