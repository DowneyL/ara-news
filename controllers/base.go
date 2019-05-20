package controllers

import (
	"ara-news/boot"
	"ara-news/components/response"
	"ara-news/validators"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"gopkg.in/go-playground/validator.v9"
)

type BaseController struct {
	beego.Controller
	i18n.Locale
	response.Responser
}

type IBeforeAction interface {
	BeforeAction()
}

type IAfterAction interface {
	AfterAction()
}

func (base *BaseController) Prepare() {
	base.setLangVer()
	if app, ok := base.AppController.(IBeforeAction); ok {
		app.BeforeAction()
	}
}

func (base *BaseController) Finish() {
	if app, ok := base.AppController.(IAfterAction); ok {
		app.AfterAction()
	}
}

func (base *BaseController) setLangVer() bool {
	currLocale := boot.GetLocale()
	base.Lang = currLocale.Lang
	base.Responser.Lang = currLocale.Lang

	base.Data["Lang"] = currLocale.Lang
	base.Data["CurrLangName"] = currLocale.Name
	base.Data["RestLocale"] = boot.GetRestLocale()

	return currLocale.IsNeedRedirect
}

func (base *BaseController) returnJson() {
	base.ServeJSON()
	base.StopRun()
}

func (base *BaseController) SuccessJSON(data interface{}) {
	base.Data["json"] = base.Success(data)
	base.returnJson()
}

func (base *BaseController) ErrorJSON(code response.ErrorCode, message string, tr ...bool) {
	if beego.BConfig.RunMode != beego.DEV && code == response.QUERY_ERROR {
		code = response.SYSTEM_ERROR
		message = "system error"
		tr[0] = true
	}

	translate := len(tr) > 0 && tr[0]
	base.Data["json"] = base.Error(code, message, translate)
	base.returnJson()
}

func (base *BaseController) InvalidArgumentJSON(errors ...string) {
	base.Data["json"] = base.InvalidArgument(errors...)
	base.returnJson()
}

func (base *BaseController) SystemErrorJSON(errors ...string) {
	base.Data["json"] = base.SystemError(errors...)
	base.returnJson()
}

func (base *BaseController) QueryErrorJSON(errors ...string) {
	if beego.BConfig.RunMode != beego.DEV {
		errors = []string{}
	}
	base.Data["json"] = base.QueryError(errors...)
	base.returnJson()
}

func (base *BaseController) Valid(obj interface{}) {
	valid := boot.GetValidator()
	if err := valid.Struct(obj); err != nil {
		errs := err.(validator.ValidationErrors)
		trans := validators.GetTrans(boot.GetLang())
		errTrs := errs.Translate(trans)
		var errors []string
		for _, v := range errTrs {
			errors = append(errors, v)
		}
		base.InvalidArgumentJSON(errors...)
	}
}

func (base *BaseController) ValidForm(obj interface{}) {
	if err := base.ParseForm(obj); err != nil {
		if beego.BConfig.RunMode == beego.DEV {
			base.InvalidArgumentJSON(err.Error())
		}
		base.InvalidArgumentJSON()
	}
	base.Valid(obj)
}

func (base *BaseController) ValidJSON(obj interface{}) {
	if err := json.Unmarshal(base.Ctx.Input.RequestBody, obj); err != nil {
		if beego.BConfig.RunMode == beego.DEV {
			base.InvalidArgumentJSON(err.Error())
		}
		base.InvalidArgumentJSON()
	}
	base.Valid(obj)
}
