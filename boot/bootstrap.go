package boot

import (
	"ara/components/geoip"
	"ara/components/lang"
	"ara/components/mysql"
	"ara/validators"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var App Application

type Application struct {
	GeoIP      geoip.Geo
	Locale     lang.Locale
	RestLocale []lang.Locale
	Validator  validators.UniversalValidator
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

	// 初始化基础验证器
	App.Validator = validators.InitValidate()

	// 注册 maxmind 地址库
	geoip.RegisterCityIpReader()
}

func GetGeoIP() geoip.Geo {
	return App.GeoIP
}

func GetValidator() validators.UniversalValidator {
	return App.Validator
}

func GetLocale() lang.Locale {
	return App.Locale
}

func GetRestLocale() []lang.Locale {
	return App.RestLocale
}
