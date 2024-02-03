package models

type Agent struct {
	BaseModel
	ID       int    `json:"id"`
	Type     int    `json:"type"`
	Pid      string `json:"pid"` // Parent ID
	Virtual  int    `json:"virtual"`
	UserName string `json:"user_name"`
	Balance  string `json:"balance"`
	NickName string `json:"nick_name"`
	Status   int    `json:"status"`
}
