package models

type UserBalance struct {
	BaseModel
	Recharge     string `json:"recharge"`
	Withdraw     string `json:"withdraw"`
	ProfitOrLost string `json:"ks"`
	Balance      string `json:"balance"`
}
