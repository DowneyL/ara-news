package main

import (
	_ "ara-news/boot"
	_ "ara-news/routers"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.SetStaticPath("/assets", "images")

	beego.Run()
}
