package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["ipaas/controllers/app:AppController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:AppController"],
		beego.ControllerComments{
			Method: "CreateApp",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:AppController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:AppController"],
		beego.ControllerComments{
			Method: "ListApp",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:AppController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:AppController"],
		beego.ControllerComments{
			Method: "DeleteApp",
			Router: `/:app`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:AppController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:AppController"],
		beego.ControllerComments{
			Method: "OperationApp",
			Router: `/:app/:verb`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ConfigController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ConfigController"],
		beego.ControllerComments{
			Method: "CreateConfig",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ConfigController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ConfigController"],
		beego.ControllerComments{
			Method: "ListConfig",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ConfigController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ConfigController"],
		beego.ControllerComments{
			Method: "DeleteConfig",
			Router: `/:config`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ConfigController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ConfigController"],
		beego.ControllerComments{
			Method: "AddConfigData",
			Router: `/:config`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ConfigController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ConfigController"],
		beego.ControllerComments{
			Method: "DeleteConfigData",
			Router: `/:config/:key`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ContainerController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ContainerController"],
		beego.ControllerComments{
			Method: "ListContainer",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ContainerController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ContainerController"],
		beego.ControllerComments{
			Method: "ReCreateContainer",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ContainerController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ContainerController"],
		beego.ControllerComments{
			Method: "GetConatainerMetric",
			Router: `/:name/metrics`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"],
		beego.ControllerComments{
			Method: "CreateService",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"],
		beego.ControllerComments{
			Method: "ListService",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"],
		beego.ControllerComments{
			Method: "DeleteService",
			Router: `/:service`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"],
		beego.ControllerComments{
			Method: "OperatorService",
			Router: `/:service/:verb`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"],
		beego.ControllerComments{
			Method: "GetOperation",
			Router: `/:service/audits`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"],
		beego.ControllerComments{
			Method: "ListServiceEvents",
			Router: `/:service/events`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:ServiceController"],
		beego.ControllerComments{
			Method: "GetServiceMetric",
			Router: `/:service/metrics`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:StorageController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:StorageController"],
		beego.ControllerComments{
			Method: "CreateStorage",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:StorageController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:StorageController"],
		beego.ControllerComments{
			Method: "ListStorage",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/app:StorageController"] = append(beego.GlobalControllerRouter["ipaas/controllers/app:StorageController"],
		beego.ControllerComments{
			Method: "DeleteStorage",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
