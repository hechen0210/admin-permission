package request

import (
	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
)

type AdminRequest struct {
}

type AdminForm struct {
	Id       int    `json:"id"`
	Account  string `json:"account" validate:"required|minLen:6|maxLen:20"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required|minLen:6|maxLen:20"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Group    int    `json:"group" validate:"required"`
	Status   int    `json:"status"`
}

func NewAdminRequest() *AdminRequest {
	return &AdminRequest{}
}

func (ar *AdminRequest) GetFormData(ctx iris.Context) (data AdminForm, err error) {
	err = ctx.ReadJSON(&data)
	if err != nil {
		return
	}
	validate := validate.Struct(data)
	return data, validate.ValidateE().OneError()
}
