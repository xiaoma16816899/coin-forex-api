package models

// 资金报表
type FundStatement struct {
	Date             string  `json:"date"`
	Recharge         string  `json:"recharge"`
	RechargeNums     int     `json:"rechargeNums"`
	RechargeUserNums int     `json:"rechargeUserNums"`
	Withdraw         string  `json:"withdraw"`
	WithdrawNums     int     `json:"withdrawNums"`
	WithdrawUserNums int     `json:"withdrawUserNums"`
	Kesun            float64 `json:"kesun"`
	Commission       string  `json:"commission"`
	UserNums         int     `json:"userNums"`
	ActiveNums       int     `json:"activeNums"`
}
