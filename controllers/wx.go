package controllers

import (
	"server-hub/library"
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
	return library.NewWX().GetAccessToken()
}
