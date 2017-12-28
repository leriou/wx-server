package controllers

import (
	// "fmt"
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/go-redis/redis"
)

/**
 * 微信公众号控制器
 */
type WxController struct{
	MainController
}

// 获取access_token
func (c *WxController) getAccessToken() string{
	//从redis中获取access_token
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// access_key := "wx_access_token"
	val, _ := client.Get("wx_access_token").Result()
	if val == "" {
		return c.refresh()
	} else {
		return val
	}
}

// 刷新access_token,并更新redis
func (c *WxController) refresh() string{
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

	access_key := "wx_access_token"
	timeoutDuration := time.Duration(expires_in) * time.Second
	err := client.Set(access_key, access_token, timeoutDuration).Err()
	if err != nil {
		panic(err)
	}
	return access_token
}

