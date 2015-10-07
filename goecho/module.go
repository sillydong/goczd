package goecho

type Module interface {
	InitRoute()
}

type StructResponse struct {
	Errno int         `json:errno,omitempty`
	Error string      `json:"error,omitempty"`
	Total int64       `json:total,omitempty`
	Data  interface{} `json:data,omitempty`
}
