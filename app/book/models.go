package book

/*
type BookStatus struct {
	BookId				int
	SeatId				string 		`json:"store_id"`
	BookDate			string 		`json:"store_name"`
	BookBeginTime		// 10:00		`json:"province"`
	BookEndTime			// 10:15 string		`json:"city"`
	Status			string		`json:"address"`  // valid, pending, booked, invalid
	OpenId			string		`json:"phone_no"`
}

type WxBookOrder struct {
	OrderId 			string
	WxOrderId
	BookId
	OpenId				string 		`json:"store_id"`
	TranType			// pay, refund
	TranAmount			int 		`json:"store_name"`
}*/

type BookReq struct {
	OpenId        string `json:"openid"`
	ShopId        string `json:"shop_id"`
	ShopName      string `json:"shop_name"`
	SeatId        string `json:"seat_id"`
	SeatName      string `json:"seat_name"`
	SeatType      string `json:"seat_type"`
	BookBeginDate string `json:"book_begin_date"`
	BookEndDate   string `json:"book_end_date"`
	BookBeginTime string `json:"book_begin_time"`
	BookEndTime   string `json:"book_end_time"`
	BookDuration  string `json:"book_duration"`
	OrderType     string `json:"order_type"` // 订单类型，accountSave：充值账户，bookConsume：预定位置支付
	PayMethod     string `json:"pay_method"` // 支付方式，wechat、account、giftcard 等
	Currency      string `json:"currency"`   // 支付币种，默认前端传CNY
	Amount        int64  `json:"amount"`     // 支付金额，单位：分
	PayInfo       string `json:"pay_info"`
	TranTime      int64  `json:"tran_time"` // 交易时间，unix时间戳格式
}

type BookResp struct {
	OutTradeNo     string `json:"out_trade_no"`
	PrepayId       string `json:"prepay_id"`
	ChannelOrderNo string `json:"channel_order_no"`
}

type BookOrderReq struct {
	OpenId       string `json:"openid"`
	OrderType    string `json:"order_type"`
	OutTradeNo   string `json:"out_trade_no"`
	BeginDate    string `json:"begin_date"`
	EndDate      string `json:"end_date"`
	NotifyStatus string `json:"notify_status"`
	Page         int    `json:"page"`
	PageSize     int    `json:"page_size"`
}

type BookOrderDivideByTimeReq struct {
	OpenId     string `json:"openid"`
	FutureFlag string `json:"future_flag"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
}

type BookOrderResp struct {
	OpenId         string `json:"openid"`
	ChannelOrderNo string `json:"channel_order_no"`
	OutTradeNo     string `json:"out_trade_no"`
	ShopId         string `json:"shop_id"`
	SeatId         string `json:"seat_id"`
	SeatType       string `json:"seat_type"`
	BookBeginDate  string `json:"book_begin_date"`
	BookEndDate    string `json:"book_end_date"`
	BookBeginTime  string `json:"book_begin_time"`
	BookEndTime    string `json:"book_end_time"`
	BookDuration   string `json:"book_duration"`
	OrderType      string `json:"order_type"`
	PayMethod      string `json:"pay_method"`
	Currency       string `json:"currency"`
	Amount         int64  `json:"amount"`
	PayInfo        string `json:"pay_info"`
}

type BookOrderUpdateReq struct {
	OpenId         string `json:"openid"`
	OutTradeNo     string `json:"out_trade_no"`
	ChannelOrderNo string `json:"begin_date"`
	PayStatus      string `json:"pay_status"`
	NotifyStatus   string `json:"notify_status"`
}

type BookSeatOrderStatus struct {
	StoreId       string `json:"store_id"`
	StoreName     string `json:"store_name"`
	SeatId        string `json:"seat_id"`
	SeatName      string `json:"seat_name"`
	OpenId        string `json:"openid" gorm:"column:openid"`
	BookBeginTime string `json:"book_begin_time"`
	BookEndTime   string `json:"book_end_time"`
	Status        string `json:"status"`
}

type SeatOrderStatusReq struct {
	StoreId       string `json:"store_id"`
	StoreName     string `json:"store_name"`
	SeatId        string `json:"seat_id"`
	SeatName      string `json:"seat_name"`
	OpenId        string `json:"openid" gorm:"column:openid"`
	BookBeginTime string `json:"book_begin_time"`
	BookEndTime   string `json:"book_end_time"`
	ToStatus      string `json:"to_status"`
}
