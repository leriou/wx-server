package controllers

import (
	"encoding/xml"
	"time"
	"server-hub/models"
)

/*******************************
 * controllers 所可能用到的结构体*
 *******************************/

/**
 * 返回结果 {"status":false,"msg":"request error"}
 */
type apiResponse struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Data   []models.User
}

/**
 * xml请求结构体
 */
type XMLBody struct {
	XMLName	xml.Name `xml:"xml"`
	ToUserName	string	 `xml:ToUserName`
	FromUserName	string	`xml:FromUserName`
	CreateTime	time.Duration	`xml:CreateTime`
	MsgType	string	`xml:MsgType`
	Content	string	`xml:Content`
	MsgId	int	`xml:MsgId`
}

/**
 * 用户响应的XML结构体
 */
type XMLResponse struct {
	XMLName	xml.Name `xml:"xml"`
	ToUserName	string `xml:ToUserName`
	FromUserName	string	`xml:FromUserName`
	CreateTime	time.Duration	`xml:CreateTime`
	MsgType	string	`xml:MsgType`
	Content	string	`xml:Content`
}
/**
 * 微信事件
 */
type XMLWxEvent struct {
	XMLName xml.Name `xml:"xml`
	ToUserName	string `xml:ToUserName`
	FromUserName	string	`xml:FromUserName`
	CreateTime	time.Duration `xml.CreateTime`
	Event string 	`xml:Event`
	MsgType	string	`xml:MsgType`
}