package billing

type InsertBillingReq struct {
	OutTradeNo     string `json:"out_trade_no"`
	ChannelOrderNo string `json:"channel_order_no"`
	OrderType      string `json:"order_type"` // 订单类型，accountSave：充值账户，bookConsume：预定位置支付
	PayMethod      string `json:"pay_method"`
	OpenId         string `json:"openid" gorm:"column:openid"`
	StoreId        string `json:"store_id"`
	StoreName      string `json:"store_name"`
	SeatId         string `json:"seat_id"`
	SeatName       string `json:"seat_name"`
	SeatType       string `json:"seat_type"`
	BookBeginDate  string `json:"book_begin_date"`
	BookEndDate    string `json:"book_end_date"`
	BookBeginTime  string `json:"book_begin_time"`
	BookEndTime    string `json:"book_end_time"`
	BookDuration   string `json:"book_duration"`
	Currency       string `json:"currency"`
	Amount         int64  `json:"amount"` // 金额，单位：分
	PayInfo        string `json:"pay_info"`
	TranTime       int64  `json:"tran_time"` // 交易时间，unixtime时间戳格式
	Remark         string `json:"remark"`
	BillStatus     string `json:"bill_status"`   // 订单状态：VALID：生效， INVALID：无效
	PayStatus      string `json:"pay_status"`    // 支付状态，同步返回，INIT：初始状态，SUCCESS：支付成功， PENDING：支付处理中，FAIL：支付失败
	NotifyStatus   string `json:"notify_status"` // 支付异步返回状态，INIT：初始状态，SUCCESS：支付成功， PENDING：支付处理中，FAIL：支付失败
}

type InsertBillingResp struct {
	OutTradeNo string `json:"out_trade_no"`
}

type BillingInfoReq struct {
	OpenId       string `json:"openid"`
	OrderType    string `json:"order_type"`
	OutTradeNo   string `json:"out_trade_no"`
	BeginDate    string `json:"begin_date"`
	EndDate      string `json:"end_date"`
	NotifyStatus string `json:"notify_status"`
	Page         int    `json:"page"`
	PageSize     int    `json:"page_size"`
}

type BillingInfoDivideByTimeReq struct {
	OpenId     string `json:"openid"`
	FutureFlag string `json:"future_flag"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
}

type BillingInfoResp struct {
	OpenId         string `json:"openid" gorm:"column:openid"`
	ChannelOrderNo string `json:"channel_order_no"`
	OutTradeNo     string `json:"out_trade_no"`
	ShopId         string `json:"shop_id" gorm:"column:store_id"`
	ShopName       string `json:"shop_name" gorm:"column:store_name"`
	SeatId         string `json:"seat_id"`
	SeatName       string `json:"seat_name"`
	SeatType       string `json:"seat_type"`
	BookBeginDate  string `json:"book_begin_date"`
	BookEndDate    string `json:"book_end_date"`
	BookBeginTime  string `json:"book_begin_time"`
	BookEndTime    string `json:"book_end_time"`
	BookDuration   string `json:"book_duration"`
	BeginUnixTime  int64  `json:"begin_unix_time"`
	EndUnixTime    int64  `json:"end_unix_time"`
	OrderType      string `json:"order_type"`
	PayMethod      string `json:"pay_method"`
	Currency       string `json:"currency"`
	Amount         int64  `json:"amount"`
	TranTime       int64  `json:"tran_time"`
	PayInfo        string `json:"pay_info"`
	PayStatus      string `json:"pay_status"`
}

type UpdateBillingReq struct {
	OpenId         string `json:"openid"`
	OutTradeNo     string `json:"out_trade_no"`
	ChannelOrderNo string `json:"channel_order_no"`
	PayStatus      string `json:"pay_status"`
	NotifyStatus   string `json:"notify_status"`
}

type OrderDetail struct {
	OutTradeNo     string `json:"out_trade_no"`
	ChannelOrderNo string `json:"channel_order_no"`
	OrderType      string `json:"order_type"` // 订单类型，accountSave：充值账户，bookConsume：预定位置支付
	PayMethod      string `json:"pay_method"`
	OpenId         string `json:"openid" gorm:"column:openid"`
	StoreId        string `json:"store_id"`
	StoreName      string `json:"store_name"`
	SeatId         string `json:"seat_id"`
	SeatName       string `json:"seat_name"`
	SeatType       string `json:"seat_type"`
	BookBeginDate  string `json:"book_begin_date"`
	BookEndDate    string `json:"book_end_date"`
	BookBeginTime  string `json:"book_begin_time"`
	BookEndTime    string `json:"book_end_time"`
	BookDuration   string `json:"book_duration"`
	BeginUnixTime  int64  `json:"begin_unix_time"`
	EndUnixTime    int64  `json:"end_unix_time"`
	Currency       string `json:"currency"`
	Amount         int64  `json:"amount"` // 金额，单位：分
	PayInfo        string `json:"pay_info"`
	TranTime       int64  `json:"tran_time"` // 交易时间，unixtime时间戳格式
	Remark         string `json:"remark"`
	BillStatus     string `json:"bill_status"`   // 订单状态：VALID：生效， INVALID：无效
	PayStatus      string `json:"pay_status"`    // 支付状态，同步返回，Init：初始状态，Success：支付成功， Pending：支付处理中，FAIL：支付失败
	NotifyStatus   string `json:"notify_status"` // 支付异步返回状态，Init：初始状态，Success：支付成功， Pending：支付处理中，FAIL：支付失败
}

type BillingStatusCheckReq struct {
	CheckUnixTime int `json:"check_unix_time"`
}

type BillingStatusCheckResp struct {
	CheckUnixTime  int    `json:"check_unix_time"`
	StoreId        string `json:"store_id"`
	SeatId         string `json:"seat_id"`
	OutTradeNo     string `json:"out_trade_no"`
	ChannelOrderNo string `json:"channel_order_no"`
	BeginUnixTime  int    `json:"begin_unix_time"`
	EndUnixTime    int    `json:"end_unix_time"`
}
