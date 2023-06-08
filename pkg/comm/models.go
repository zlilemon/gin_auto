package comm

// Result http return common result
type Result struct {
	Code    int         `json:"errcode"`
	Message string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

const AppName = "gss_server"
