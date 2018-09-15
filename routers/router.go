package routers

import (
	"github.com/astaxie/beego"
	"github.com/hlca/base_api/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/token", &controllers.AuthController{})
}
