package account

import (
	"time"
)

type ChargeReq struct {
	OpenId         string `json:"openid"`
	OutTradeNO     string `json:"out_trade_no"`     // 内部订单号，传到充值记录中，方便查询
	ChannelOrderNo string `json:"channel_order_no"` // 渠道订单号，传到充值记录中，方便查询
	Currency       string `json:"currency"`         // 充值币种，默认前端透传 CNY
	Amount         int64  `json:"amount"`           // 充值金额，单位：分
	PayMethod      string `json:"pay_method"`       // 支付方式：wechat、alipay……
	//ChargeTime     string `json:"charge_time"`      // YYYY-MM-DD HH:MM:SS 格式
	ChargeTime int64  `json:"charge_time"` // unix_time 格式
	PayInfo    string `json:"pay_info"`    // 交易附言，kv格式拼接
	VersionId  int64  `json:"version_id"`
}

type ChargeResp struct {
	OutTradeNo     string `json:"out_trade_no"`     // 内部订单号
	ChannelOrderNo string `json:"channel_order_no"` // 支付渠道侧订单号
	OpenId         string `json:"openid"`
	Currency       string `json:"currency"`
	Amount         int64  `json:"amount"` // 充值金额，单位：分
}

type BalanceReq struct {
	OpenId string `json:"openid"`
}

type BalanceResp struct {
	OpenId     string `json:"openid"`
	Currency   string `json:"currency"`
	Amount     int64  `json:"amount"`
	Status     string `json:"status"`
	VersionId  int64  `json:"version_id"`
	UpdateTime int64  `json:"updated_at"`
}

type BalanceModel struct {
	OpenId     string `json:"openid"`
	Currency   string `json:"currency"`
	Amount     int64  `json:"amount"`
	Status     string `json:"status"`
	VersionId  int64  `json:"version_id"`
	UpdateTime int64  `json:"updated_at"`
}

type WaterModel struct {
	StaticDate       string    `json:"static_date"`
	OpenId           string    `json:"openid" gorm:"column:openid"`
	OutTradeNo       string    `json:"out_trade_no"`
	ChannelOrderNo   string    `json:"channel_order_no"`
	TranType         string    `json:"tran_type"`
	SubTranType      string    `json:"sub_tran_type"`
	Currency         string    `json:"currency"`
	Amount           int64     `json:"amount"` // 金额，单位：分
	TranTime         string    `json:"tran_time"`
	TranTimeUnixTime int       `json:"tran_time_unix_time"` // 时间戳格式
	Remark           string    `json:"remark"`
	CreateAt         time.Time `json:"create_at"`
	UpdateAt         time.Time `json:"update_at"`
}

type BalanceSaveReq struct {
	OpenId    string `json:"openid"`
	Currency  string `json:"currency"`
	Amount    int64  `json:"amount"`
	VersionId int64  `json:"version_id"`
}

type BalanceSaveResp struct {
	OpenId           string `json:"openid"`
	Currency         string `json:"currency"`
	BeforeSaveAmount int64  `json:"before_amount"`
	AfterSaveAmount  int64  `json:"after_amount"`
	VersionId        int64  `json:"version_id"`
}

type ConsumeReq struct {
	OpenId     string `json:"openid"`
	OutTradeNO string `json:"out_trade_no"` // 内部订单号
	//ChannelOrderNo	string `json:"channel_order_no"`	// 渠道订单号
	Currency string `json:"currency"` // 充值币种，默认前端透传 CNY
	Amount   int64  `json:"amount"`   // 充值金额，单位：分
	//PayMethod  		string `json:"pay_method"`			// 支付方式：wechat、alipay……
	ConsumeTime int64 `json:"consume_time"` // unix_time 时间戳格式
	//ConsumeTime 	LocalTime 			`json:"consume_time"` // YYYY-MM-DD HH:MM:SS 格式
	PayInfo   string `json:"pay_info"` // 交易附言，kv格式拼接
	VersionId int64  `json:"version_id"`
}

type ConsumeResp struct {
	OutTradeNo string `json:"out_trade_no"` // 内部订单号
	OpenId     string `json:"openid"`
	Currency   string `json:"currency"`
	Amount     int64  `json:"amount"` // 充值金额，单位：分
}

type WaterListReq struct {
	OpenId    string `json:"openid"`
	TranType  string `json:"tran_type"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

type WaterListResp struct {
	WaterList []*WaterInfo `json:"water_list"`
}

type WaterInfo struct {
	TranTime       string `json:"tran_time"` // 交易时间：格式：YYYY-MM-DD HH:MM:SS
	OpenId         string `json:"openid"`
	TranType       string `json:"tran_type"` // 交易类型，charge：充值， consume：消耗
	Currency       string `json:"currency"`
	Amount         int64  `json:"amount"` // 充值金额，单位：分
	OutTradeNo     string `json:"out_trade_no"`
	ChannelOrderNo string `json:"channel_order_no"`
	PayMethod      string `json:"pay_method"` // 支付方式
}

type WaterDetailReq struct {
	OutTradeNo string `json:"out_trade_no"` // 内部订单号
}

type GetChargeGearReq struct {
	Type         string `json:"type"`
	ChargeAmount int64  `json:"charge_amount"`
}

type GetChargeGearResp struct {
	ChargeGearList []ChargeGearItem `json:"charge_gear_list"`
}

type ChargeGearItem struct {
	RelationId   string `json:"relation_id"`
	RelationName string `json:"relation_name"`
	ChargeAmount int64  `json:"charge_amount"`
	AccountPoint int64  `json:"account_point"`
}

type ChargeGearInfo struct {
	RelationId   string `json:"relation_id"`
	RelationName string `json:"relation_name"`
	ChargeAmount int64  `json:"charge_amount"`
	AccountPoint int64  `json:"account_point"`
}
