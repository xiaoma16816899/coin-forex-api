package models

/*
	Status 全部, 正常, 禁用
*/
// 用户钱包
type UserWallet struct {
	BaseModel
	UserID   int    `json:"user_id"`
	Currency string `json:"currency"`
	Icon     string `json:"icon"`
	Chain    string `json:"chain"`
	Address  string `json:"address"`
	Status   int    `json:"status"`
}

type UserWalletDetail struct {
	UserWallet
	User      UserInfo `json:"user"`
	AgentList []Agent  `json:"agent_list"`
}
