package models

type ProductConfig struct {
	BaseModel
	ProductTypeID  int    `json:"product_type_id"`
	Title          string `json:"title"`
	Name           string `json:"name"`
	Icon           string `json:"icon"`
	Price          string `json:"price"`
	FloatPrice     string `json:"float_price"`
	Bd             int    `json:"bd"`
	Fixed          int    `json:"fixed"`
	AmountFixed    int    `json:"amount_fixed"`
	Kline          string `json:"kline"`
	API            int    `json:"api"`
	APIConfig      string `json:"api_config"`
	TimeConfig     string `json:"time_config"`
	ContractConfig string `json:"contract_config"`
	TimeSet        string `json:"time_set"`
	FeeRate        string `json:"fee_rate"`
	Min            int    `json:"min"`
	Introduce      string `json:"introduce"`
	Buy            string `json:"buy"`
	Hot            int    `json:"hot"`
	Sort           int    `json:"sort"`
	Status         int    `json:"status"`
}
