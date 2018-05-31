package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["ipaas/controllers/system:NodeController"] = append(beego.GlobalControllerRouter["ipaas/controllers/system:NodeController"],
		beego.ControllerComments{
			Method: "ListNode",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/system:NodeController"] = append(beego.GlobalControllerRouter["ipaas/controllers/system:NodeController"],
		beego.ControllerComments{
			Method: "GetNode",
			Router: `/:name`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/system:NodeController"] = append(beego.GlobalControllerRouter["ipaas/controllers/system:NodeController"],
		beego.ControllerComments{
			Method: "GetNodeMetric",
			Router: `/:name/metrics`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/system:NodeController"] = append(beego.GlobalControllerRouter["ipaas/controllers/system:NodeController"],
		beego.ControllerComments{
			Method: "Scheduler",
			Router: `/:node`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/system:NodeController"] = append(beego.GlobalControllerRouter["ipaas/controllers/system:NodeController"],
		beego.ControllerComments{
			Method: "ListContainer",
			Router: `/:node/containers`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ipaas/controllers/system:NodeController"] = append(beego.GlobalControllerRouter["ipaas/controllers/system:NodeController"],
		beego.ControllerComments{
			Method: "LabelOperator",
			Router: `/:node/labels/:verb`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

}
