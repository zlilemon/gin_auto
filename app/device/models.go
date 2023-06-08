package device

type GetDeviceInfoReq struct {
}

type GetDeviceInfoResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrorCode   int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

type DeviceMapping struct {
	StoreId            string `json:"store_id"`             // 设备所在等店铺id
	SeatId             string `json:"seat_id"`              // 座位id
	DeviceFunctionType string `json:"device_function_type"` // 设备作用分类：entrance_door：入户门， switch：开关
	DeviceId           string `json:"device_id"`            // 硬件设备对应等设备id
	DeviceName         string `json:"device_name"`
	DeviceCategory     string `json:"device_category"` // 开关、电灯、led 等
	DeviceBrand        string `json:"device_brand"`    // 设备品牌，eg：tuya
	Status             string `json:"status"`          // 设备状态，100：有效； 500：失效；
	remark             string `json:"remark"`
}

type OperationReq struct {
	OutTradeNo         string `json:"out_trade_no"`
	OpenId             string `json:"openid"`
	DeviceFunctionType string `json:"device_function_type"`
	Cmd                string `json:"cmd"`
	StoreId            string `json:"store_id"`
	SeatId             string `json:"seat_id"`
	DeviceCategory     string `json:"device_category"`
	DeviceId           string `json:"device_id"`
	DeviceBrand        string `json:"device_brand"`
}

type OperationResp struct {
}

type TuyaReq struct {
	Commands []TuyaCommands `json:"commands"`
}

type TuyaCommands struct {
	Code  string `json:"code"`
	Value bool   `json:"value"`
}

type XiaoMiReq struct {
	EntityId string `json:"entity_id"`
}
