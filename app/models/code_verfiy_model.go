package models

// 发送记录
type VerifyCode struct {
	BaseModel
	Account   string `json:"account"`
	Content   string `json:"content"`
	Code      string `json:"code"`
	IP        string `json:"ip"`
	Reason    string `json:"reason"`
	Status    int    `json:"status"`
	IPAddress string `json:"ipAddress"`
	TelecomName string `json:"telecom_name"`
}
