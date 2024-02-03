package models

type ContractRisk []struct {
	BaseModel
	UserID   int    `json:"user_id"`
	Risk     int    `json:"risk"`
	Operator string `json:"operator"`
	Status   int    `json:"status"`
}

type ContractRiskDetail struct {
	ContractRisk
	UserInfo  UserInfo `json:"user"`
	AgentList []Agent  `json:"agentList"`
}
