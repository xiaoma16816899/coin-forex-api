package models

// 用户IP记录 (全段)
type UserIPAddress []struct {
	BaseModel
	UserID   int    `json:"user_id"`
	Title    string `json:"title"`
	Platform string `json:"platform"`
	IP       string `json:"ip"`
	Host     string `json:"host"`
	Path     string `json:"path"`
	Reason   string `json:"reason"`
	Status   int    `json:"status"`
	Address  string `json:"address"`
	Isp      string `json:"isp"`
}

type UserIPAddressDetail struct {
	UserIPAddress
	AgentList []Agent `json:"agent_list"`
}
