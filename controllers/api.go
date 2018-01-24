package controllers

import (
	// "server-hub/library"
)

/**
 * API中心控制器
 */
type ApiController struct{
	MainController
}

func (c *ApiController) TestJson(){
	res := apiResponse{true,"api request success"}
	c.Data["json"] = &res
	c.ServeJSON()
}
