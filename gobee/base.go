package gobee

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

const (
	_ = iota
	STATUS_OK
	STATUS_BADREQUEST
	STATUS_UNAUTHORIZED
	STATUS_NOTFOUND
	STATUS_NOMORE
	STATUS_ERROR
)

const (
	MSG_BADREQUEST   = "请求错误"
	MSG_UNAUTHORIZED = "未授权的请求"
	MSG_NOTFOUND     = "无结果"
	MSG_NOMORE       = "无更多结果"
	MSG_ERROR        = "请求出错"
)
