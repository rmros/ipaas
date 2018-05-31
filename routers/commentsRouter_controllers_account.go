package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["ipaas/controllers/account:SpaceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:SpaceController"],
		beego.ControllerComments{
			Method: "CreateSpace",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/account:SpaceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:SpaceController"],
		beego.ControllerComments{
			Method: "ListSpace",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/account:SpaceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:SpaceController"],
		beego.ControllerComments{
			Method: "DeleteSpace",
			Router: `/:namespace`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/account:SpaceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:SpaceController"],
		beego.ControllerComments{
			Method: "GetSpace",
			Router: `/:namespace`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/account:TeamController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:TeamController"],
		beego.ControllerComments{
			Method: "CreateTeam",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/account:TeamController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:TeamController"],
		beego.ControllerComments{
			Method: "ListTeam",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/account:TeamController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:TeamController"],
		beego.ControllerComments{
			Method: "DeleteTeam",
			Router: `/:team`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/account:TeamController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:TeamController"],
		beego.ControllerComments{
			Method: "AddSpace",
			Router: `/:team/spaces`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/account:TeamController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:TeamController"],
		beego.ControllerComments{
			Method: "AddUsers",
			Router: `/:team/users`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/account:TeamController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:TeamController"],
		beego.ControllerComments{
			Method: "GetUsers",
			Router: `/:team/users`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/account:UserController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:UserController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/account:UserController"] = append(beego.GlobalControllerRouter["ipaas/controllers/account:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
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
