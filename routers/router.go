package routers

import (
	"base_api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/token", &controllers.AuthController{})
}
