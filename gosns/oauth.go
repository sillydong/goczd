package gosns

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
)

const (
	WEIBO = 1 + iota
	QZONE
	WEIXIN
)

func NewOauth(t, key, secret, callback string) (*Oauth, error) {
	switch t {
	case WEIBO:
		return &Weibo{key: key, Secret: secret, Callback: callback}, nil
	case QZONE:
		return &Qzone{key: key, Secret: secret, Callback: callback}, nil
	case WEIXIN:
		return &Weixin{Key: key, Secret: secret, Callback: callback}, nil
	default:
		return nil, fmt.Errorf("%v not support", t)
	}
}

type Oauth interface {
	GetLoginUrl() string
	HandleCallback(code, baseurl string) (*AccessToken, error)
	GetUserInfo(accesstoken, openid string) (UserInfo, error)
}

type AccessToken struct {
	AccessToken string
	ExpiresIn   int64
	OpenId      string
}

type UserInfo struct {
	OpenId string
	Name   string
	Avatar string
	Gender string
}

func get(url string, v interface{}) error {
	req := httplib.Get(url)
	return req.ToJson(v)
}

func gets(url string) (string, error) {
	req := httplib.Get(url)
	return req.String()
}

func post(url string, params map[string]string, v interface{}) error {
	req := httplib.Post(url)
	if len(params) > 0 {
		for key, val := range params {
			req.Param(key, val)
		}
	}
	return req.ToJson(v)
}
