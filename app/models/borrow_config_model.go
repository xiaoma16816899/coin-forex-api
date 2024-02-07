package models

// 借款配置
type BorrowingConfig struct {
	BaseModel
	Title  string `json:"title"`
	Day    int    `json:"day"`
	Rate   string `json:"rate"`
	Sort   int    `json:"sort"`
	Status int    `json:"status"`
}
