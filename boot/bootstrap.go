package boot

import (
	"ara-news/components/geoip"
	"ara-news/components/lang"
	"ara-news/components/mysql"
	"ara-news/components/pagination"
	"ara-news/validators"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"gopkg.in/go-playground/validator.v9"
)

var App Application

type Application struct {
	GeoIP      geoip.Geo
	Locale     lang.Locale
	RestLocale []lang.Locale
	Validate   *validator.Validate
	Pagination pagination.Pagination
}

func init() {
	if beego.BConfig.RunMode != "dev" {
		err := logs.SetLogger(beego.AppConfig.String("logsAdapterFile"), `{"filename" : "runtime/runtime.log"}`)
		if err != nil {
			panic("[Error] Set logger Failed," + err.Error())
		}
	}

	// 注册所有支持的多语言文件
	lang.RegisterLangConf()

	// 初始化数据库信息
	mysql.InitMysql()

	// 初始国际化验证器
	App.Validate = validators.InitValidate()

	// 注册 maxmind 地址库
	geoip.RegisterCityIpReader()
}

func GetGeoIP() geoip.Geo {
	return App.GeoIP
}

func GetValidator() *validator.Validate {
	return App.Validate
}

func GetLocale() lang.Locale {
	return App.Locale
}

func GetRestLocale() []lang.Locale {
	return App.RestLocale
}

func GetPagination() pagination.Pagination {
	return App.Pagination
}

func GetLang() string {
	return GetLocale().Lang
}
