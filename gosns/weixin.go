package gosns
import (
	"fmt"
	"math/rand"
	"github.com/cosiner/ygo/resource"
	"github.com/Unknwon/com"
	"net/url"
)
type Weixin struct {
	Key      string
	Secret   string
	Callback string
}

func (oauth *Weixin)GetLoginUrl() string {
	return fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?&response_type=code&scope=snsapi_userinfo&state=%d&appid=%s&redirect_url=%s",rand.Int(),oauth.Key,url.QueryEscape(oauth.Callback))
}

type WeixinAccessToken struct{
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	OpenId      string `json:"openid"`
}

func (oauth *Weixin)HandleCallback(code string) (*AccessToken, error) {
	taccess := &WeixinAccessToken{}
	err:=get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",oauth.Key,oauth.Secret,code),taccess)
	if err != nil {
		return nil,err
	}else{
		return &AccessToken{AccessToken:taccess.AccessToken,ExpiresIn:taccess.ExpiresIn,OpenId:taccess.OpenId},nil
	}
}

type WeixinUserInfo struct{
	OpenId     string `json:"openid"`
	NickName   string `json:"nickname"`
	HeadImgUrl string `json:"headimgurl"`
}

func (oauth *Weixin)GetUserInfo(accesstoken, openid string) (*UserInfo, error) {
	tuserinfo := &WeixinUserInfo{}
	err:=get(fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s",accesstoken,openid),&tuserinfo)
	if err != nil {
		return nil,err
	}else{
		return &UserInfo{OpenId:tuserinfo.OpenId,Name:tuserinfo.NickName,Avatar:tuserinfo.HeadImgUrl}
	}
}
