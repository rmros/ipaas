package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["ipaas/controllers/app:AppController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:AppController"],
		beego.ControllerComments{
			Method: "CreateApp",
			Router: `/apps`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:AppController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:AppController"],
		beego.ControllerComments{
			Method: "ListApp",
			Router: `/apps`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:AppController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:AppController"],
		beego.ControllerComments{
			Method: "DeleteApp",
			Router: `/apps/:app`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:AppController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:AppController"],
		beego.ControllerComments{
			Method: "ReDeployApp",
			Router: `/apps/:app/redeploy`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:AppController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:AppController"],
		beego.ControllerComments{
			Method: "StartApp",
			Router: `/apps/:app/start`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:AppController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:AppController"],
		beego.ControllerComments{
			Method: "StopApp",
			Router: `/apps/:app/stop`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"],
		beego.ControllerComments{
			Method: "CreateService",
			Router: `/services`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"],
		beego.ControllerComments{
			Method: "ListService",
			Router: `/services`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"],
		beego.ControllerComments{
			Method: "ReDeployService",
			Router: `/services/:service/redeploy`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"],
		beego.ControllerComments{
			Method: "StartService",
			Router: `/services/:service/start`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"],
		beego.ControllerComments{
			Method: "ReStartService",
			Router: `/services/:service/start`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"],
		beego.ControllerComments{
			Method: "DeleteService",
			Router: `/services/service`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"],
		beego.ControllerComments{
			Method: "StopService",
			Router: `/services/service/stop`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

}
