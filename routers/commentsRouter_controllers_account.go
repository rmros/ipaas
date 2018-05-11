package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["ipaas/controllers/account:UserController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/account:UserController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
