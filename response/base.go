package response

import (
	"admin-permission/config"

	"github.com/kataras/iris"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) SetCode(code int) *Response {
	r.Code = code
	return r
}

func (r *Response) SetMessage(message string) *Response {
	r.Message = message
	return r
}

func (r *Response) SetData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) Error(ctx iris.Context) {
	r.Code = config.REQUEST_ERROR
	r.Message = config.CODE_TITLE[r.Code]
	ctx.JSON(r)
}

func (r *Response) Fail(ctx iris.Context) {
	if r.Code == 0 {
		r.Code = config.REQUEST_FAIL
	}
	if r.Message == "" {
		r.Message = config.CODE_TITLE[r.Code]
	}
	ctx.JSON(r)
}

func (r *Response) Success(ctx iris.Context) {
	if r.Code == 0 {
		r.Code = config.REQUEST_SUCCESS
	}
	if r.Message == "" {
		r.Message = config.CODE_TITLE[r.Code]
	}
	ctx.JSON(r)
}
