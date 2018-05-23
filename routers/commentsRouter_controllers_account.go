package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["ipaas/controllers/account:UserController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:UserController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/account:UserController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:UserController"],
		beego.ControllerComments{
			Method: "ResetPassword",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/account:UserController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:UserController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/account:UserController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:user`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

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
			Router: `/logout`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
