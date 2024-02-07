package models

/*
	    type: 0 彩金 = bonus
	    type: 1 = recharge
		type: 2 = Withdraw
		type: 3 = Commission
		type: 4 = Order = 订单
		type: 8 = borrow
		type: 10 = unknown
		Status: 1 = show data on frontend when use check on fund details
		Status: 0 = hide data on frontend when use check on fund detail no see
*/

// 账变记录
type FundManagementRecord struct {
	BaseModel
	UserID   int    `json:"user_id"`
	Mode     int    `json:"mode"`
	Type     int    `json:"type"`
	Title    string `json:"title"`
	Des      string `json:"des"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
	Balance  string `json:"balance"`
	KeyID    int    `json:"key_id"`
	Remark   string `json:"remark"`
	Status   int    `json:"status"`
}

type FundManagementLogDetail struct {
	FundManagementRecord
	UserInfo  UserInformation `json:"user"`
	AgentList []Agent         `json:"agent_list"`
}
