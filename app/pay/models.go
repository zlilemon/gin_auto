package pay

import (
	"encoding/xml"
)

type WxPayReq struct {
	OpenId     string `json:"openid"`
	Amount     int64  `json:"amount"`
	PayInfo    string `json:"pay_info"`
	OutTradeNo string `json:"out_trade_no"`
}

type WxPayResp struct {
	PrepayId  string `json:"prepay_id"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

type QueryReq struct {
	TransactionId string `json:"transaction_id"`
	MchId         string `json:"mchid"`
}

type QueryResp struct {
}

type WxPayReqParam struct {
	WxAppid    string `schema:"wx_appid"       validate:"nonzero,max=30"`
	TradeType  string `schema:"trade_type"     validate:"nonzero,max=9"`
	OutTradeNo string `schema:"out_trade_no"   validate:"nonzero,max=30"`
	Amount     int64  `schema:"amount"         validate:"nonzero"`
	WxOpenId   string `schema:"wx_openid"      validate:"nonzero,max=30"`
	PayInfo    string `schema:"pay_info"       validate:"max=256"`
	//	Ts         int    `schema:"ts"             validate:"nonzero"`
	//	Sign       string `schema:"sign"           validate:"nonzero,max=32"`
}

type WxPayResponse struct {
	PrepayId string `json:"prepay_id"`
}

type WxpayNotifyReq struct {
	ID           string              `json:"id"`
	CreateTime   string              `json:"create_time"`
	ResourceType string              `json:"resource_type"`
	EvenType     string              `json:"even_type"`
	Summary      string              `json:"summary"`
	Resource     WxpayNotifyResource `json:"resource"`
}

type WxpayNotifyResource struct {
	OriginalType   string `json:"original_type"`
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Nonce          string `json:"nonce"`
}

type WxJsapiAmount struct {
	Total    int64  `json:"total"`
	Currency string `json:"currency"`
}

type WxJsapiPayer struct {
	Openid string `json:"openid"`
}

type WxJsapiOrderReq struct {
	XMLName xml.Name `xml:"xml,omitempty"`
	AppId   string   `xml:"appid,omitempty"`
	//Body       string   `xml:"body,omitempty"`
	MchId       string `xml:"mchid,omitempty"`
	Description string `xml:"description,omitempty"`
	OutTradeNo  string `xml:"out_trade_no,omitempty"`
	Attach      string `xml:"attach,omitempty"`
	NotifyUrl   string `xml:"notify_url,omitempty"`
	Amount      WxJsapiAmount
	Payer       WxJsapiPayer
	NonceStr    string `xml:"nonce_str,omitempty"`
}

type WxJsapiOrderRsp struct {
	XMLName    xml.Name `xml:"xml,omitempty"`
	ReturnCode string   `xml:"return_code,omitempty"`
	ReturnMsg  string   `xml:"return_msg,omitempty"`
	ResultCode string   `xml:"result_code,omitempty"`
	ErrCode    string   `xml:"err_code,omitempty"`
	ErrCodeDes string   `xml:"err_code_des,omitempty"`
	PrepayId   string   `xml:"prepay_id,omitempty"`
}
