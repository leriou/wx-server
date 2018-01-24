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


func (c *UserController) Create(){
	u := models.User{}
	u.Username = "test1"
	u.Age = 8
	u.State = 1
	u.Mobile = "17621111877"
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

func (c *UserController) Update() {
	u := models.User{Id:1}
	u.Username = "test2"
	u.State = 2
	models.UpdateUser(&u)
	c.EnableRender = false
}

func (c *UserController) Remove() {
	models.DeleteUser(3)
	c.EnableRender = false
}

