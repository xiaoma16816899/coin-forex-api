package model

type ForexTradingModel struct {
	ID                   uint64  `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	ClosePrice           float64 `json:"close_price"`
	OpeningPrice         float64 `json:"opening_price"`
	HighPrice            float64 `json:"high_price"`
	LowPrice             float64 `json:"low_price"`
	TransactionsVolume   float64 `json:"transactions_volume"`
	OTC                  float32 `json:"otc"`
	TradingTime          float64 `json:"trading_time"`
	TradingVolume        float64 `json:"trading_volume"`
	WeightedAveragePrice float64 `json:"weighted_average_price"`
}

type ParamForexTradingModel struct {
	CurrencyType string `json:"currency_type"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	Limit        int    `json:"limit"`
}
