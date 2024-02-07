package models

type Option struct {
	BaseModel
	UserID     int     `json:"user_id"`
	ProductID  int     `json:"product_id"`
	OrderNo    string  `json:"order_no"`
	Title      string  `json:"title"`
	Fixed      int     `json:"fixed"`
	Side       int     `json:"side"`
	Time       int     `json:"time"`
	Profit     string  `json:"profit"`
	Amount     float64 `json:"amount"`
	Fee        float64 `json:"fee"`
	OpenPrice  float64 `json:"open_price"`
	SellPrice  float64 `json:"sell_price"`
	SellTime   int     `json:"sell_time"`
	SellProfit float64 `json:"sell_profit"`
	RiskStatus int     `json:"risk_status"`
	Operator   string  `json:"operator"`
	Status     int     `json:"status"`
}

type OptionDetail struct {
	Option
	UserInfo  UserInformation `json:"user"`
	AgentList []Agent         `json:"agent_list"`
}
