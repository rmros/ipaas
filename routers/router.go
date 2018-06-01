// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ipaas/controllers/account"
	"ipaas/controllers/app"
	"ipaas/controllers/system"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/users",
			beego.NSInclude(
				&account.UserController{},
			),
		),
		beego.NSNamespace("/teams",
			beego.NSInclude(
				&account.TeamController{},
			),
		),
		beego.NSNamespace("/clusters/:cluster/nodes",
			beego.NSInclude(
				&system.NodeController{},
			),
		),
		beego.NSNamespace("/clusters/:cluster",
			beego.NSInclude(
				&system.ClusterController{},
			),
		),
		beego.NSNamespace("/clusters/:cluster/namespaces/:namespace/apps",
			beego.NSInclude(
				&app.AppController{},
			),
		),
		beego.NSNamespace("/clusters/:cluster/namespaces/:namespace/services",
			beego.NSInclude(
				&app.ServiceController{},
			),
		),
		beego.NSNamespace("/clusters/:cluster/namespaces/:namespace/storages",
			beego.NSInclude(
				&app.StorageController{},
			),
		),
		beego.NSNamespace("/clusters/:cluster/namespaces/:namespace/configs",
			beego.NSInclude(
				&app.ConfigController{},
			),
		),
		beego.NSNamespace("/clusters/:cluster/namespaces/:namespace/containers",
			beego.NSInclude(
				&app.ContainerController{},
			),
		),
		beego.NSNamespace("/clusters/:cluster/nodes",
			beego.NSInclude(
				&system.NodeController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
