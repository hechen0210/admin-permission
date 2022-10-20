package controller

import (
	"admin-permission/request"
	"admin-permission/response"
	"admin-permission/service"

	"github.com/kataras/iris/v12"
)

type LoginController struct {
	request  *request.LoginRequest
	service  *service.AdminService
	response *response.Response
}

func NewLoginController() *LoginController {
	return &LoginController{
		request:  request.NewLoginRequest(),
		service:  service.NewAdminService(),
		response: response.NewResponse(),
	}
}

func (lc *LoginController) Login(ctx iris.Context) {
	data, err := lc.request.GetLoginForm(ctx)
	if err != nil {
		lc.response.Error(ctx)
	}
	lc.service.Login(data.Account, data.Password)
}

func (lc *LoginController) Logout(ctx iris.Context) {
}
