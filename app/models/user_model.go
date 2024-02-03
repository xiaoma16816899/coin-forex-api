package models

/*
Status: 全部, 正常, 禁用, 禁止提现, 禁止下单, 禁止提现和下单, 禁止团队返佣
*/
// 用户列表
type User struct {
	BaseModel
	Type              int     `json:"type"`
	Virtual           int     `json:"virtual"`
	RoleID            int     `json:"role_id"`
	LevelID           int     `json:"level_id"`
	Layer             int     `json:"layer"`
	Pid               string  `json:"pid"`
	Tid               string  `json:"tid"`
	UserName          string  `json:"user_name"`
	NickName          string  `json:"nick_name"`
	Password          string  `json:"_"`
	FundPassword      string  `json:"-"`
	Avatar            string  `json:"avatar"`
	AreaCode          int     `json:"area_code"`
	Mobile            string  `json:"mobile"`
	Email             string  `json:"email"`
	Balance           string  `json:"balance"`
	LoanAmount        string  `json:"loan_amount"`
	Secret            string  `json:"secret"`
	Auth              string  `json:"auth"`
	VisitorID         string  `json:"visitor_id"`
	RechargeAmount    string  `json:"recharge_amount"`
	Remark            string  `json:"remark"`
	Credit            int     `json:"credit"`
	JSON              string  `json:"json"` // store use address
	CreateIP          string  `json:"create_ip"`
	LoginIP           string  `json:"login_ip"`
	LoginTime         int     `json:"login_time"`
	Status            int     `json:"status"`
	Yk                int     `json:"yk"`
	VisitorIDNums     int     `json:"visitor_id_nums"`
	CreateIPNums      int     `json:"create_ip_nums"`
	LoginIPNums       int     `json:"login_ip_nums"`
	IPAddress         string  `json:"ipAddress"`
	IsOnline          bool    `json:"isOnline"`
	RebateAmount      string  `json:"rebate_amount"`
	TodayRebateAmount int     `json:"today_rebate_amount"`
	Recharge          string  `json:"recharge"`
	RechargeNums      int     `json:"recharge_nums"`
	Withdraw          string  `json:"withdraw"`
	WithdrawNums      int     `json:"withdraw_nums"`
	ProfitOrLost      float64 `json:"ks"`
	TeamMember        Team    `json:"team"`
	RegNums           int     `json:"regNums"`
}

type UserInfo struct {
	ID       int    `json:"id"`
	Type     int    `json:"type"`
	Pid      string `json:"pid"`
	Layer    int    `json:"layer"`
	UserName string `json:"user_name"`
	Virtual  int    `json:"virtual"`
	Status   int    `json:"status"`
}
