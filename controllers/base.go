package controllers

import (
	// "github.com/go-redis/redis"
	"github.com/astaxie/beego"
	"server-hub/library"
	"time"
	"io/ioutil"
	"encoding/xml"
	"strings"
	"github.com/go-redis/redis"
)

/**
 * 默认控制器
 */
type MainController struct {
	beego.Controller
}

func init(){

}

func (c *MainController) GetRedis() *redis.Client{
	di := new(library.Di)
	return di.GetRedis()
}

/**
 * 默认Get请求
 * @author lixiumeng@gongchang.com
 * @addtime 2017-07-12T01:38:54+0800
 * @version 1.0.0
 * @param   {[type]}                 c *MainController) Get( [description]
 * @return  {[type]}                   [description]
 */
func (c *MainController) Get() {
	var echostr string
	echostr = c.GetString("echostr") //随机字符串
	if echostr == "" {
		echostr = "request success"
	} else{
		// 获取四个微信参数,进行字典排序
		token := "serverhub"
		signature := c.GetString("signature") // 签名
		timestamp := c.GetString("timestamp") // 时间戳
		nonce := c.GetString("nonce")         // 随机数
		//验证微信服务器的加密
		WX := library.WX{}
		if !WX.ValidateWxServer(token,signature,timestamp,nonce) {
			echostr = ""
		}
	}
	// 返回微信echostr,接入成功
	c.Ctx.WriteString(echostr)
}

/**
 * 微信post过来的xml或者json信息
 * @author lixiumeng@gongchang.com
 * @addtime 2017-07-12T01:39:11+0800
 * @version 1.0.0
 * @param   {[type]}                 c *MainController) Post( [description]
 * @return  {[type]}                   [description]
 */
func (c *MainController) Post() {
	var s string = ""
	if bytes,err := ioutil.ReadAll(c.Ctx.Request.Body);err == nil {
		wxXmlBody := new(XMLBody)
		xml.Unmarshal(bytes, wxXmlBody)
		s = c.dealWithXml(wxXmlBody)
	}
	c.Ctx.WriteString(s)
	c.EnableRender = false
}
/**
 * 处理接受的xml
 */
func (c *MainController) dealWithXml(xmlData *XMLBody) string{
	responseStr := c.getResponseMsg(xmlData.Content,xmlData.FromUserName)
	responseTextBody,err := c.makeResponseBody(xmlData.ToUserName,xmlData.FromUserName,responseStr)
	if err != nil {
		return ""
	}
	return string(responseTextBody)
}

/**
 * 构造返回的xml信息
 * @author lixiumeng@gongchang.com
 * @addtime 2017-07-12T01:39:21+0800
 * @version 1.0.0
 * @param   {[type]}                 c *MainController) makeResponseBody(FromUserName,ToUserName,Content string) ([]byte,error [description]
 * @return  {[type]}                   [description]
 */
func (c *MainController) makeResponseBody(FromUserName,ToUserName,Content string) ([]byte,error){
	textResponseBody := &XMLResponse{}
    textResponseBody.FromUserName = FromUserName
    textResponseBody.ToUserName = ToUserName
    textResponseBody.MsgType = "text"
    textResponseBody.Content = Content
    textResponseBody.CreateTime = time.Duration(time.Now().Unix())
    return xml.MarshalIndent(textResponseBody, " ", "  ")
}
/**
 * 根据发送信息生成返回信息
 */
func (c *MainController) getResponseMsg(sent,openid string) string{
	var s string
	v := strings.Split(sent, ":")
	tools := library.Tools{}
	// WX := library.WX{}
	switch v[0] {
		case "功能列表":
			s = `
				1 每日一句
				2 诗号查询
				3 游戏推荐
				4 音乐推荐
				5 软件推荐
				6 md5加密(md5:你好)
				7 base64加密/解密(base64_encode:你好/base64_decode:asp=)
				8 sha1加密
				8 我最喜欢谁
				1~5 功能请直接输入汉字使用  如: 每日一句
			`
		case "md5":
			s = tools.Md5(v[1])
		case "base64_encode":
			s = tools.Base64_encode(v[1])
		case "base64_decode":
			s = tools.Base64_decode(v[1])
		case "sha1":
			s = tools.Sha1(v[1])
		case "每日一句":
			s = "There is no small act of kindness. Every compassionate act makes large the world."
		case "诗号查询":
			s = `
			无情葬月:
			芳菲阑珊，
			夙缘鶗鴃，
			风驷云轩愁誓约。
			夜蝶飞阶，
			霎微雨阙，
			剑锋无情人葬月。
			`
		case "游戏推荐":
			s = `
			journey(风之旅人)
			The Last Guardian(最后的守护者)
			`
		case "音乐推荐":
			s = `
			周杰伦--夜的第七章
			周杰伦  --以父之名
			霹雳英雄--剑者传说
			霹雳英雄--荒人邪影
			霹雳英雄--夜雨寄北
			霹雳英雄--日出峨眉
			`
		case "软件推荐":
			s = `
			chrome  ---浏览器
			spotify ---国外的在线音乐软件
			medis   ---图形化的redis管理软件
			`
		default:
			client := c.GetRedis()
			val, _ := client.Get(sent).Result()
			if val == "" {
				s = "暂时不提供该服务,请稍等几天"
			} else {
				s = val
			}
	}
	return s
}
