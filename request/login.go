package request

import (
	"github.com/gookit/validate"
	"github.com/kataras/iris/v12"
)

type LoginRequest struct {
}

type LoginForm struct {
	LoginType string `json:"login_type"`
	Account   string `json:"account" validate:"required|minLen:6|maxLen:20"`
	Password  string `json:"password" validate:"required|minLen:6|maxLen:20"`
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

func (lr *LoginRequest) GetLoginForm(ctx iris.Context) (data LoginForm, err error) {
	err = ctx.ReadJSON(&data)
	if err != nil {
		return
	}
	validate := validate.Struct(data)
	return data, validate.ValidateE().OneError()
}
