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

/**
 * json返回测试
 * @author lixiumeng@gongchang.com
 * @addtime 2017-07-05T23:38:04+0800
 * @version 1.0.0
 * @param   {[type]}                 c *ApiController) TestJson( [description]
 * @return  {[type]}                   [description]
 */
func (c *ApiController) TestJson(){
	res := apiResponse{true,"api request success"}
	c.Data["json"] = &res
	c.ServeJSON()
}
