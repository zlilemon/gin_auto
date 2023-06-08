package store

type GetStoreInfoListReq struct {
	City string `json:"city"`
}

type GetStoreInfoListResp struct {
	ShopId    string  `json:"shop_id"`
	ShopName  string  `json:"shop_name"`
	Location  string  `json:"location"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	ShopPic   string  `json:"shop_pic"`
}

type StoreDetailReq struct {
	StoreId string `json:"store_id"`
}

/*
type SeatInfo struct {
	SeatId			string 		`json:"seat_id"`
	SeatName		string		`json:"seat_name"`
	BookBeginTime	string		`json:"book_begin_time"`
	BookEndTime		string		`json:"book_end_time"`
}
*/

type SeatInfoItem struct {
	SeatId       string     `json:"seat_id"`
	SeatName     string     `json:"seat_name"`
	BookTimeList []BookTime `json:"book_time_list"`
}

type SeatInfo struct {
	Normal []SeatInfoItem `json:"normal"`
	Vip    []SeatInfoItem `json:"vip"`
	Double []SeatInfoItem `json:"double"`
}

type BookTime struct {
	BookBeginTime string `json:"book_begin_time"`
	BookEndTime   string `json:"book_end_time"`
}

type StoreDetailResp struct {
	ShopId       string    `json:"shop_id"`
	ShopName     string    `json:"shop_name"`
	Location     string    `json:"location"`
	Longitude    float64   `json:"longitude"`
	Latitude     float64   `json:"latitude"`
	ShopPic      []string  `json:"shop_pic"`
	ShopDesc     string    `json:"shop_desc"`
	ShopDescPics []string  `json:"shop_desc_pics"`
	ShopSeatPics string    `json:"shop_seat_pics"`
	ShopSeats    SeatInfo  `json:"shop_seats"`
	Price        PriceInfo `json:"price"`
}

type StoreInfo struct {
	StoreId      string `json:"store_id"`
	StoreName    string `json:"store_name"`
	Province     string `json:"province"`
	City         string `json:"city"`
	Address      string `json:"address"`
	PhoneNo      string `json:"phone_no"`
	Introduction string `json:"introduction"`
	// -- 去掉 TotalSeatNum	int			`json:"total_seat_num"`
	Longitude          float64 `json:"longitude"`            // 经度
	Latitude           float64 `json:"latitude"`             // 维度
	HomePicUrl         string  `json:"home_pic_url"`         // 首页店铺图片地址
	DetailPicUrl       string  `json:"detail_pic_url"`       // 店铺详细介绍图片地址  数组
	SeatPicUrl         string  `json:"seat_pic_url"`         // 店铺 座位图片 地址
	IntroductionPicUrl string  `json:"introduction_pic_url"` // 店铺 介绍 图片 数组
}

type PriceInfo struct {
	Normal PriceItem `json:"normal"`
	Vip    PriceItem `json:"vip"`
	Double PriceItem `json:"double"`
}

type PriceItem struct {
	Minute SalePrice `json:"minute"`
	Day    SalePrice `json:"day"`
	Month  SalePrice `json:"month"`
}

type SalePrice struct {
	OrigPrice int `json:"orig"`
	RealPrice int `json:"real"`
}

type PriceModel struct {
	StoreId   string `json:"store_id"`
	StoreName string `json:"store_name"`
	SeatType  string `json:"seat_type"`
	PriceType string `json:"price_type"`
	OriPrice  int    `json:"ori_price"`
	RealPrice int    `json:"real_price"`
	Status    string `json:"status"`
}

type TbSeatInfo struct {
	StoreId   string `json:"store_id"`   // 店铺id
	StoreName string `json:"store_name"` // 店铺名称
	SeatId    string `json:"seat_id"`    // 座位id
	SeatName  string `json:"seat_name"`
	SeatType  string `json:"seat_type"` // 座位类型：vip/normal/double
}

type TbSeatOrderStatus struct {
	StoreId       string `json:"store_id"`
	SeatId        string `json:"seat_id"`
	OpenId        string `json:"openid"`          // 预定用户对应的openid
	BookBeginTime string `json:"book_begin_time"` // eg 2022-10-01 10:00
	BookEndTime   string `json:"book_end_time"`   // eg 2022-10-01 11:00
	Status        string `json:"status"`          // 数据状态： VALID：有效， INVALID：失效
}

type StoreNoticeModel struct {
	WifiName      string `json:"wifi_name"`
	WifiPasswd    string `json:"wifi_passwd"`
	CustomerPhone string `json:"customer_phone"`
}

type QuestionModel struct {
	Ask    string `json:"ask"`
	Answer string `json:"answer"`
}

type StoreNoticeResp struct {
	WifiName      string           `json:"wifi_name"`
	WifiPasswd    string           `json:"wifi_passwd"`
	CustomerPhone string           `json:"customer_phone"`
	QuestionList  []*QuestionModel `json:"question_list"`
}

type Bulletin struct {
	//BulletinId				string 	`json:"bulletin_id"`
	Id    int64  `json:"id"`
	Title string `json:"title"`
	//PublishTime 			string 	`json:"publish_time"`
	PublishTimeUnixTime int64  `json:"publish_time_unix_time"`
	PublishDetail       string `json:"publish_detail"`
	Status              string `json:"status"`
}

type BulletinSummary struct {
	//BulletinId				string 	`json:"bulletin_id"`
	Id    int64  `json:"id"`
	Title string `json:"title"`
	//PublishTime 			string 	`json:"publish_time"`
	PublishTimeUnixTime int64  `json:"publish_time_unix_time"`
	Status              string `json:"status"`
}

type BulletinSummaryResp struct {
	BulletinList []*BulletinSummary `json:"bulletin_list"`
}

type BulletinDetailReq struct {
	Id int64 `json:"id"`
}

type BulletinDetailResp struct {
	Id                  int64  `json:"id"`
	Title               string `json:"title"`
	PublishTimeUnixTime int64  `json:"publish_time_unix_time"`
	PublishDetail       string `json:"publish_detail"`
	Status              string `json:"status"`
}
