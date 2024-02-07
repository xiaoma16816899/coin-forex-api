package models

// 秒合约风控
type ContractRisk []struct {
	BaseModel
	UserID   int    `json:"user_id"`
	Risk     int    `json:"risk"`
	Operator string `json:"operator"`
	Status   int    `json:"status"`
}

type ContractRiskDetail struct {
	ContractRisk
	UserInfo  UserInformation `json:"user"`
	AgentList []Agent         `json:"agentList"`
}
