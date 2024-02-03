package models

/*
status number 0 is 禁用 ( Disable )
status number 1 is 正常 ( Enable )
status empty "" mean that 全部 ( All )
*/
type PaymentAddress struct {
	BaseModel
	Address  string
	Chain    string
	Currency string // USDT, Other
	Icon     string
	Status   int
}
