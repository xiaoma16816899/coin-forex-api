package models

// 快捷调价
type Adjustment struct {
	BaseModel
	Title  string  `json:"title"`
	Fixed  int     `json:"fixed"`
	Price  string  `json:"price"`
	API    int     `json:"api"`
	Bodong float64 `json:"bodong"`
}

type AdjustmentDetail struct {
	Adjustment
	APIConfig
}

type APIConfig struct {
	Symbol string `json:"symbol"`
	Ratio  string `json:"ratio"`
	Min    string `json:"min"`
	Max    string `json:"max"`
	MinVol string `json:"minVol"`
	MaxVol string `json:"maxVol"`
}
