package models

// 合同订单
type Contract struct {
	BaseModel
	UserID       int    `json:"user_id"`
	ProductID    int    `json:"product_id"`
	OrderNo      string `json:"order_no"`
	Title        string `json:"title"`
	Side         int    `json:"side"`
	Fixed        int    `json:"fixed"`
	Multiple     int    `json:"multiple"`
	Amount       string `json:"amount"`
	Fee          string `json:"fee"`
	HandNums     string `json:"hand_nums"`
	ContractNums string `json:"contract_nums"`
	TpPrice      string `json:"tp_price"`
	SlPrice      string `json:"sl_price"`
	WtPrice      string `json:"wt_price"`
	OpenPrice    string `json:"open_price"`
	OpenTime     int    `json:"open_time"`
	SellPrice    string `json:"sell_price"`
	SellTime     int    `json:"sell_time"`
	SellProfit   string `json:"sell_profit"`
	EnforcePrice string `json:"enforce_price"`
	Status       int    `json:"status"`
}

type ContractDetail struct {
	Contract
	UserInfo    UserInformation `json:"user"`
	AgentList   []Agent         `json:"agent_list"`
	ProductInfo Product         `json:"product"`
}

type Product struct {
	ID    int    `json:"id"`
	Price string `json:"price"`
}
