package controllers

import (
	"ara/boot"
	"ara/components/response"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"gopkg.in/go-playground/validator.v9"
	"strings"
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
	translate := len(tr) > 0 && tr[0]
	base.Data["json"] = base.Error(code, message, translate)
	base.returnJson()
}

func (base *BaseController) InvalidArgumentJSON() {
	base.Data["json"] = base.InvalidArgument()
	base.returnJson()
}

func (base *BaseController) SystemErrorJSON() {
	base.Data["json"] = base.SystemError()
	base.returnJson()
}

func (base *BaseController) Valid(obj interface{}) {
	valid := boot.GetValidator()
	if err := valid.Validate.Struct(obj); err != nil {
		errs := err.(validator.ValidationErrors)
		errTrs := errs.Translate(valid.Trans)
		var errstr []string
		for _, v := range errTrs {
			errstr = append(errstr, v)
		}
		base.ErrorJSON(response.PARAMS_ERROR, strings.Join(errstr, "\r\n"))
	}
}

func (base *BaseController) ValidForm(obj interface{}) {
	if err := base.ParseForm(obj); err != nil {
		base.InvalidArgumentJSON()
	}
	base.Valid(obj)
}

func (base *BaseController) ValidJSON(obj interface{}) {
	if err := json.Unmarshal(base.Ctx.Input.RequestBody, obj); err != nil {
		base.InvalidArgumentJSON()
	}
	base.Valid(obj)
}
