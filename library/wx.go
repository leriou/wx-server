package library

import(
	"sort"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	simplejson "github.com/bitly/go-simplejson"
)

type WX struct {

}

/**
 * 验证微信服务器
 */
func (w *WX) ValidateWxServer(wx_token,wx_signature,wx_timestamp,wx_nonce string) bool{
	tmps := []string{wx_token, wx_timestamp, wx_nonce}
	sort.Strings(tmps)
	tmpStr := tmps[0] + tmps[1] + tmps[2]
	tools := Tools{}
	tmp := tools.Sha1(tmpStr)
	return tmp == wx_signature
}
// get_access_token
func (w *WX) GetAccessToken() (string,int){
	//从微信接口获取新的access_token
	token_url :=  beego.AppConfig.String("wx.token_url")
	grant_type := beego.AppConfig.String("wx.grant_type")
	appid := beego.AppConfig.String("wx.appid")
	secret := beego.AppConfig.String("wx.secret")
	wx_access_token := token_url +"?grant_type="+grant_type+"&appid=" +appid + "&secret="+secret
	s,_ := httplib.Get(wx_access_token).String()
	//从json中提取所需信息
	js,_ := simplejson.NewJson([]byte(s))
	access_token,_ := js.Get("access_token").String() // access_token
	expires_in,_ := js.Get("expires_in").Int()  // default 7200
	return access_token,expires_in
}





