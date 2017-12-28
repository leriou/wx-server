package routers

import (
	"github.com/astaxie/beego"
	"server-hub/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    // registe auto router
    beego.AutoRouter(&controllers.UserController{})
    beego.AutoRouter(&controllers.WxController{})
    beego.AutoRouter(&controllers.ApiController{}) 
}
