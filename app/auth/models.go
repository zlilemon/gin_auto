package auth

type FreshTokenReq struct {
}

type FreshTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrorCode   int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type Code2SessionRequest struct {
	Appid     string `query:"appid"`
	Secret    string `query:"secret"`
	JsCode    string `query:"js_code"`
	GrantType string `query:"grant_type"`
}

type Code2SessionResponse struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	ErrorCode  int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrorCode   int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

type AccessTokenInfo struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
