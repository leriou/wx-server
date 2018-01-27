package controllers

import (
	"server-hub/models"
)

/**
 * API中心控制器
 */
type ApiController struct{
	MainController
}

func (c *ApiController) TestJson(){
	c.Data["json"] = &apiResponse{true,"api request success",[]models.User{}}
	c.ServeJSON()
}
