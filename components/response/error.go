package response

import (
	"ara-news/helper"
	"github.com/beego/i18n"
)

type Responser struct {
	Response
	i18n.Locale
}

func (r *Responser) Success(data interface{}) Response {
	r.Status.Code = SUCCESS
	r.Status.Message = r.Tr("success")
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
	r.Status.Time = helper.Date("Y-m-d H:i:s")
	r.Data = new(struct{})

	return r.Response
}

func (res *Response) parseErrors(errors ...string) {
	if len(errors) > 0 {
		res.Status.Errors = errors
	}
}

func (r *Responser) InvalidArgument(errors ...string) Response {
	res := r.Error(PARAMS_ERROR, "params error", true)
	res.parseErrors(errors...)

	return res
}

func (r *Responser) SystemError(errors ...string) Response {
	res := r.Error(SYSTEM_ERROR, "system error", true)
	res.parseErrors(errors...)

	return res
}

func (r *Responser) QueryError(errors ...string) Response {
	res := r.Error(QUERY_ERROR, "query error", true)
	res.parseErrors(errors...)

	return res
}
