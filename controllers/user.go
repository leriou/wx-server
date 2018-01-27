package controllers

import (
	"server-hub/models"
	"time"
	"server-hub/library"
	"github.com/m1ome/randstr"
)

/**
 * 用户功能控制器
 */
type UserController struct{
	ApiController
}

func (c *UserController) Create(){
	u := models.User{}
	u.Username = c.GetString("username")
	u.Age,_ = c.GetInt64("age")
	u.State = 1
	u.Mobile = c.GetString("mobile")
	u.Email = c.GetString("email")
	u.Salt = randstr.GetString(6)
	u.Password = library.NewTools().Md5(c.GetString("password")+u.Salt)
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
	id,_ := c.GetInt64("id")
	_,u := models.FindUserById(id)
	u.Username = c.GetString("username")
	u.State = 1
	models.UpdateUser(&u)
	c.EnableRender = false
	c.Ctx.WriteString("success")
}

func (c *UserController) Remove() {
	id,_ := c.GetInt64("id")
	models.DeleteUser(id)
	c.EnableRender = false
	c.Ctx.WriteString("success")
}

func (c *UserController) GetUser() {
	id,_ := c.GetInt64("id")
	_,u := models.FindUserById(id)
	c.Data["json"] = &apiResponse{true,"success",[]models.User{u}}
	c.ServeJSON()
}

