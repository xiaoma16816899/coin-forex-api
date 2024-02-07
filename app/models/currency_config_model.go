package models

/*
Type = 1 ( Recharge )
Type = 2 ( Withdraw )
*/
type CurrencyConfig struct {
	BaseModel
	Type   int         `json:"type"`
	Mode   int         `json:"mode"`
	Title  string      `json:"title"`
	Icon   string      `json:"icon"`
	IsBind int         `json:"is_bind"`
	Info   interface{} `json:"info"`
	Fn     string      `json:"fn"`
	Min    int         `json:"min"`
	Max    int         `json:"max"`
	Remark string      `json:"remark"`
	Sort   int         `json:"sort"`
	Status int         `json:"status"`
}
