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
	"ara-news/controllers"
	"ara-news/controllers/news"
	"ara-news/validators"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, func(context *context.Context) {
		// 获取当前语言
		boot.App.Locale, boot.App.RestLocale = lang.InitLang(context)

		boot.App.GeoIP = geoip.InitGeoIP(context)

		boot.App.Validator = validators.InitUniversalValidator(boot.App.Locale.Lang)

		boot.App.Pagination = pagination.InitPagination(context)
	})

	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/hello", &controllers.HelloController{}),
		beego.NSRouter("/user", &controllers.UserController{}, "get:GetAll;post:Create"),
		beego.NSRouter("/user/:id([0-9]+)", &controllers.UserController{}),

		beego.NSNamespace("/news",
			beego.NSRouter("/", &news.Controller{}, "post:Create"),
			beego.NSRouter("/:id([0-9]+)", &news.Controller{}, "get:Detail"),
			beego.NSNamespace("/category",
				beego.NSRouter("/", &news.CategoryController{}, "get:List;post:Create;delete:BatchDelete"),
				beego.NSRouter("/:id([0-9]+)", &news.CategoryController{}, "get:Detail;delete:Delete;put:Update"),
				beego.NSRouter("/:id([0-9]+)/name-en", &news.CategoryController{}, "patch:UpdateNameEn"),
			),
		),
	)

	beego.AddNamespace(ns)
}
