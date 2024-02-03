package models

// 首页
type ReportDetail struct {
	OnlineNums                    int     `json:"online_nums"`
	UserBalance                   string  `json:"user_balance"`
	ReviewUserWithdrawAmount      int     `json:"review_user_withdraw_amount"`
	TodayReviewWithdrawAmount     string  `json:"today_review_withdraw_amount"`
	TotalQuantifyUserNums         int     `json:"total_quantify_user_nums"`
	YesterdayQuantifyUserNums     int     `json:"yesterday_quantify_user_nums"`
	DateQuantifyUserNums          int     `json:"dateQuantifyUserNums"`
	TotalRechargeAmount           string  `json:"totalRechargeAmount"`
	TotalRechargeNums             int     `json:"totalRechargeNums"`
	TotalRechargeUserNums         int     `json:"totalRechargeUserNums"`
	TotalWithdrawAmount           string  `json:"totalWithdrawAmount"`
	TotalWithdrawNums             int     `json:"totalWithdrawNums"`
	TotalWithdrawUserNums         int     `json:"totalWithdrawUserNums"`
	Totalkesun                    float64 `json:"totalkesun"` // Take the deposit amount subtract with the withdraw amount
	DateRechargeAmount            string  `json:"dateRechargeAmount"`
	DateRechargeNums              int     `json:"dateRechargeNums"`
	DateRechargeUserNums          int     `json:"dateRechargeUserNums"`
	DateWithdrawAmount            string  `json:"dateWithdrawAmount"`
	DateWithdrawNums              int     `json:"dateWithdrawNums"`
	DateWithdrawUserNums          int     `json:"dateWithdrawUserNums"`
	YesterdayRechargeAmount       int     `json:"yesterdayRechargeAmount"`
	YesterdayRechargeNums         int     `json:"yesterdayRechargeNums"`
	YesterdayRechargeUserNums     int     `json:"yesterdayRechargeUserNums"`
	YesterdayWithdrawAmount       int     `json:"yesterdayWithdrawAmount"`
	YesterdayWithdrawNums         int     `json:"yesterdayWithdrawNums"`
	YesterdayWithdrawUserNums     int     `json:"yesterdayWithdrawUserNums"`
	Yesterdaykesun                int     `json:"yesterdaykesun"`
	TotalFirstRechargeNums        int     `json:"totalFirstRechargeNums"`
	YesterdayFirstRechargeNums    int     `json:"yesterdayFirstRechargeNums"`
	DateFirstRechargeNums         int     `json:"dateFirstRechargeNums"`
	TotalRepeatRechargeAmount     int     `json:"totalRepeatRechargeAmount"`
	TotalRepeatRechargeNums       int     `json:"totalRepeatRechargeNums"`
	YesterdayRepeatRechargeAmount int     `json:"yesterdayRepeatRechargeAmount"`
	YesterdayRepeatRechargeNums   int     `json:"yesterdayRepeatRechargeNums"`
	DateRepeatRechargeAmount      int     `json:"dateRepeatRechargeAmount"`
	DateRepeatRechargeNums        int     `json:"dateRepeatRechargeNums"`
	Totalcaijin                   int     `json:"totalcaijin"` // total amount that we are recharge from system
	Yesterdaycaijin               int     `json:"yesterdaycaijin"`
	Datecaijin                    int     `json:"datecaijin"`
	TotalRegNums                  int     `json:"totalRegNums"`
	DateRegNums                   int     `json:"dateRegNums"`
	YesterdayRegNums              int     `json:"yesterdayRegNums"`
}
