// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ara-news/boot"
	"ara-news/components/geoip"
	"ara-news/components/lang"
	"ara-news/components/pagination"
	"ara-news/controllers/help"
	"ara-news/controllers/news"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, func(context *context.Context) {
		// 获取当前语言
		boot.App.Locale, boot.App.RestLocale = lang.InitLang(context)

		boot.App.GeoIP = geoip.InitGeoIP(context)

		boot.App.Pagination = pagination.InitPagination(context)
	})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/news",
			beego.NSRouter("/", &news.Controller{}, "get:List;post:Create"),
			beego.NSRouter("/:id([0-9]+)", &news.Controller{}, "get:Detail;delete:Delete"),
			beego.NSRouter("/:id([0-9]+)/content", &news.Controller{}, "post:CreateContent"),
			beego.NSNamespace("/category",
				beego.NSRouter("/", &news.CategoryController{}, "get:List;post:Create;delete:BatchDelete"),
				beego.NSRouter("/:id([0-9]+)", &news.CategoryController{}, "get:Detail;delete:Delete;put:Update"),
				beego.NSRouter("/:id([0-9]+)/name-en", &news.CategoryController{}, "patch:UpdateNameEn"),
			),
		),
		beego.NSNamespace("/help",
			beego.NSNamespace("/category",
				beego.NSRouter("/", &help.CategoryController{}, "post:Create"),
			),
		),
	)

	beego.AddNamespace(ns)
}
