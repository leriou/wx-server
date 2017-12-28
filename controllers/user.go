package controllers

import (
	"server-hub/models"
	"time"
)

/**
 * 用户功能控制器
 */
type UserController struct{
	ApiController
}

/**
 * 用户创建
 * @author lixiumeng@gongchang.com
 * @addtime 2017-07-05T23:40:33+0800
 * @version 1.0.0
 * @param   {[type]}                 c *UserController) Create( [description]
 * @return  {[type]}                   [description]
 */
func (c *UserController) Create(){
	u := models.User{}
	u.Username = "test1"
	u.Age = 8
	u.State = 1
	u.Mobile = "17621191057"
	u.Email = "yange1@qq.com"
	u.Salt = "5xgt"
	u.Password = "kakahhhhuu11"
	u.Created = time.Now().String()
	u.Updated = time.Now().String()
	id := models.SaveUser(&u)
	if id > 0 {
		c.Ctx.WriteString("success")
	} else {
		c.Ctx.WriteString("error")
	}
}

/**
 * 用户更新
 * @author lixiumeng@gongchang.com
 * @addtime 2017-07-13T23:49:28+0800
 * @version 1.0.0
 * @param   {[type]}                 c *UserController) Update( [description]
 * @return  {[type]}                   [description]
 */
func (c *UserController) Update() {
	u := models.User{Id:1}
	u.Username = "test2"
	u.State = 2
	models.UpdateUser(&u)
	c.EnableRender = false
}

/**
 * 用户删除
 * @author lixiumeng@gongchang.com
 * @addtime 2017-07-13T23:49:38+0800
 * @version 1.0.0
 * @param   {[type]}                 c *UserController) Remove( [description]
 * @return  {[type]}                   [description]
 */
func (c *UserController) Remove() {
	models.DeleteUser(3)
	c.EnableRender = false
}

