package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["base_api/controllers:AuthController"] = append(beego.GlobalControllerRouter["base_api/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/token`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["base_api/controllers:AuthController"] = append(beego.GlobalControllerRouter["base_api/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Token",
			Router: `/token2`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
