package gohttp

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/google/go-querystring/query"
	"strings"
)

//返回数据接口
type Response interface {
	OK() bool    //返回是否正确
	Msg() string //返回数据中的错误信息
}

//请求接口
type Request interface {
	B() error             //请求发起之前执行
	URL() string          //返回请求的地址
	V() (bool, error)     //判断请求是否合格
	H() map[string]string //设置请求头
}

//打开DEBUG模式，将打印发起的请求信息
var REQUEST_DEBUG bool = false

//设置请求的UserAgent
var REQUEST_USERAGENT = ""

//发起GET请求返回struct
func DoGetResponse(req Request, resp Response) error {
	bytes, err := DoGet(req)
	if err != nil {
		return err
	} else {
		return json.Unmarshal(bytes, resp)
	}
}

//发起GET请求
func DoGet(req Request) ([]byte, error) {
	if err := req.B(); err != nil {
		return nil, err
	}
	params, err := query.Values(req)
	if err != nil {
		return nil, err
	} else {
		if ok, err := req.V(); ok {
			url := req.URL()
			if url != "" {
				if len(params) > 0 {
					if strings.Contains(url, "?") {
						url = url + "&" + params.Encode()
					} else {
						url = url + "?" + params.Encode()
					}
				}
				client := httplib.Get(url).SetEnableCookie(true)
				if REQUEST_DEBUG {
					client.Debug(true)
					client.DumpBody(true)
				}
				if REQUEST_USERAGENT != "" {
					client.SetUserAgent(REQUEST_USERAGENT)
				}
				headers := req.H()
				if len(headers) > 0 {
					for key, val := range headers {
						client.Header(key, val)
					}
				}
				bytes, err := client.Bytes()
				if err != nil {
					return nil, err
				} else {
					return bytes, nil
				}
			} else {
				return nil, fmt.Errorf("REQUEST错误: %s\n", url)
			}
		} else {
			return nil, err
		}
	}
}

//发起POST请求返回struct
func DoPostResponse(req Request, resp Response) error {
	bytes, err := DoPost(req)
	if err != nil {
		return err
	} else {
		return json.Unmarshal(bytes, resp)
	}
}

//发起POST请求
func DoPost(req Request) ([]byte, error) {
	if err := req.B(); err != nil {
		return nil, err
	}
	params, err := query.Values(req)
	if err != nil {
		return nil, err
	} else {
		if ok, err := req.V(); ok {
			url := req.URL()
			if url != "" {
				client := httplib.Post(url).SetEnableCookie(true)
				if REQUEST_DEBUG {
					client.Debug(true)
					client.DumpBody(true)
				}
				if REQUEST_USERAGENT != "" {
					client.SetUserAgent(REQUEST_USERAGENT)
				}
				if len(params) > 0 {
					for key, value := range params {
						client.Param(key, value[0])
					}
				}
				headers := req.H()
				if len(headers) > 0 {
					for key, val := range headers {
						client.Header(key, val)
					}
				}
				bytes, err := client.Bytes()
				if err != nil {
					return nil, err
				} else {
					return bytes, nil
				}
			} else {
				return nil, fmt.Errorf("REQUEST错误: %s\n", url)
			}
		} else {
			return nil, err
		}
	}
}
