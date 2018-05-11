package main

import (
	"ipaas/pkg/tools/configz"
	_ "ipaas/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	go configz.HeatLoad()
	beego.Run()
}
