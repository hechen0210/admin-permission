package controller

import (
	"admin-permission/request"

	"github.com/kataras/iris/v12"
)

type LoginController struct {
	request *request.LoginRequest
}

func NewLoginController() *LoginController {
	return &LoginController{
		request: request.NewLoginRequest(),
	}
}

func (lc *LoginController) Login(ctx iris.Context) {
	data, err := lc.request.GetLoginForm(ctx)
	if err != nil {

	}
}

func (lc *LoginController) Logout(ctx iris.Context) {
}