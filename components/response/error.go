package response

import (
	"ara/helper"
	"github.com/beego/i18n"
)

type Responser struct {
	Response
	i18n.Locale
}

func (r *Responser) Success(data interface{}) Response {
	r.Status.Code = SUCCESS
	r.Status.Message = r.Tr("success")
	r.Status.State = "keep"
	r.Status.Time = helper.Date("Y-m-d H:i:s")
	r.Data = data

	return r.Response
}

func (r *Responser) Error(err ErrorCode, message string, tr ...bool) Response {
	r.Status.Code = err
	if len(tr) > 0 && tr[0] {
		r.Status.Message = r.Tr(message)
	} else {
		r.Status.Message = message
	}
	r.Status.State = "keep"
	r.Status.Time = helper.Date("Y-m-d H:i:s")
	r.Data = new(struct{})

	return r.Response
}

func (r *Responser) InvalidArgument() Response {
	return r.Error(PARAMS_ERROR, "params error", true)
}

func (r *Responser) SystemError() Response {
	return r.Error(SYSTEM_ERROR, "system error", true)
}
