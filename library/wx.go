package library

import(
	"sort"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	simplejson "github.com/bitly/go-simplejson"
	"time"
)

type WX struct {
}

const (
	access_key = "wx_access_token"
)

func NewWX() *WX{
	return new(WX)
}
// 获取微信服务器地址
func (w *WX) GetWxServerUrl() string{
	wxConfig,_ := beego.AppConfig.GetSection("wx")
	//从微信接口获取新的access_token
	wx_server_url := wxConfig["token_url"] +"?grant_type="+wxConfig["grant_type"]+"&appid=" +wxConfig["appid"] + "&secret="+wxConfig["secret"]
	return  wx_server_url
} 
// 从微信服务器获取token
func (w *WX) GetToken() (string,int){
	s,_ := httplib.Get(w.GetWxServerUrl()).String()
	//从json中提取所需信息
	js,_ := simplejson.NewJson([]byte(s))
	access_token,_ := js.Get("access_token").String() // access_token
	expires_in,_ := js.Get("expires_in").Int()  // default 7200
	return access_token,expires_in
}
/**
 * 验证微信签名
 */
func (w *WX) ValidateWxSignature(wx_signature,wx_timestamp,wx_nonce string) bool{
	wxConfig,_ := beego.AppConfig.GetSection("wx")
	tmps := []string{wxConfig["wx_token"], wx_timestamp, wx_nonce}
	sort.Strings(tmps)
	return NewTools().Sha1(tmps[0] + tmps[1] + tmps[2]) == wx_signature
}
// get_access_token
func (w *WX) GetAccessToken() (string){
	token,err := NewDi().GetRedis().Get(access_key).Result()
	if err != nil {
		panic(err)
	}
	if token == "" {
		return w.RefreshToken()
	} else {
		return token
	}	
}
// 刷新token
func (w *WX) RefreshToken() (string){
	token,expire_time := w.GetToken()
	timeoutDuration := time.Duration(expire_time) * time.Second
	err := NewDi().GetRedis().Set(access_key,token, timeoutDuration).Err()
	if err != nil {
		panic(err)
	}
	return token
}