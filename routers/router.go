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

	"github.com/astaxie/beego"
)

func init() {
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
		beego.NSNamespace("/clusters/:cluster",
			beego.NSInclude(
				&account.SpaceController{},
			),
			beego.NSInclude(
				&app.AppController{},
			),
			beego.NSInclude(
				&app.ServiceController{},
			),
			beego.NSInclude(
				&app.StorageController{},
			),
			beego.NSInclude(
				&app.ConfigController{},
			),
			beego.NSInclude(
				&app.ContainerController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
