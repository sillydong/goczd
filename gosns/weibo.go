package gosns

import (
	"fmt"
	"net/url"
)

type Weibo struct {
	Key      string
	Secret   string
	Callback string
}

func (oauth *Weibo) GetLoginUrl() string {
	return fmt.Sprintf("https://api.weibo.com/oauth2/authorize?response_type=code&client_id=%s&redirect_url=%s", oauth.Key, url.QueryEscape(oauth.Callback))
}

type WeiboAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	OpenId      string `json:"openid"`
}

func (oauth *Weibo) HandleCallback(code string) (*AccessToken, error) {
	params := map[string]string{
		"client_id":     oauth.Key,
		"client_secret": oauth.Secret,
		"grant_type":    "authorization_code",
		"code":          code,
		"redirect_uri":  oauth.Callback,
	}
	taccess := &WeiboAccessToken{}
	err := post("https://api.weibo.com/oauth2/access_token", params, taccess)
	if err != nil {
		return nil, err
	} else {
		return &AccessToken{AccessToken: taccess.AccessToken, ExpiresIn: taccess.ExpiresIn, OpenId: taccess.OpenId}, nil
	}
}

type WeiboUserInfo struct {
	OpenId     string `json:"openid"`
	NickName   string `json:"nickname"`
	HeadImgUrl string `json:"headimgurl"`
}

func (oauth *Weibo) GetUserInfo(accesstoken, openid string) (*UserInfo, error) {
	tuserinfo := &WeiboUserInfo{}
	err := get(fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s", accesstoken, openid), tuserinfo)
	if err != nil {
		return nil, err
	} else {
		return &UserInfo{OpenId: tuserinfo.OpenId, Name: tuserinfo.NickName, Avatar: tuserinfo.HeadImgUrl}, nil
	}
}
