package user

type SaveUserInfoReq struct {
	OpenId    string `json:"openid"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	PhoneNo   string `json:"phone_no"`
}

type SaveUserInfoResp struct {
	OpenId string `json:"openid"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

type UserInfo struct {
	OpenId    string `json:"openid" gorm:"column:openid"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	PhoneNo   string `json:"phone_no"`
	UnionId   string `json:"unionid" gorm:"column:unionid"`
}

type WxPhoneResp struct {
	ErrorCode int         `json:"errcode"`
	ErrMsg    string      `json:"errmsg"`
	PhoneInfo WxPhoneInfo `json:"phone_info"`
}

type WxPhoneInfo struct {
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
}

type PhoneCode struct {
	PhoneCode string `json:"phoneCode"`
}

type ClockInfo struct {
	OpenId       string `json:"openid" gorm:"column:openid"`
	OutTradeNo   string `json:"out_trade_no"`
	StaticDate   string `json:"static_date"`
	DurationTime string `json:"duration_time"`
	Status       string `json:"status"`
}

type AddClockReq struct {
	OpenId       string `json:"openid" gorm:"column:openid"`
	OutTradeNo   string `json:"out_trade_no"`
	StaticDate   string `json:"static_date"`
	DurationTime string `json:"duration_time"`
	Status       string `json:"status"`
}

type GetClockReq struct {
	OpenId string `json:"openid" gorm:"column:openid"`
}

type AddRankingReq struct {
	OpenId       string `json:"openid" gorm:"column:openid"`
	BookTimes    int    `json:"book_times"`
	DurationTime int    `json:"duration_time"`
	NickName     string `json:"nick_name"`
	AvatarUrl    string `json:"avatar_url"`
}

type RankingInfo struct {
	BatchId             string `json:"batch_id"`
	OpenId              string `json:"openid" gorm:"column:openid"`
	WeeklyTimes         int    `json:"weekly_times"`
	WeeklyDurationTime  int    `json:"weekly_duration_time"`
	MonthlyTimes        int    `json:"monthly_times"`
	MonthlyDurationTime int    `json:"monthly_duration_time"`
	TotallyTimes        int    `json:"totally_times"`
	TotallyDurationTime int    `json:"totally_duration_time"`
}

type GetRankingListReq struct {
	RankingType string `json:"ranking_type"` //排名方式：weekly：按周， monthly：按月， totally：累计
}

type GetRankingListRespItem struct {
	OpenId       string `json:"openid"`
	Times        int    `json:"times"`
	DurationTime int    `json:"duration_time"`
	NickName     string `json:"nick_name"`
	AvatarUrl    string `json:"avatar_url"`
}
