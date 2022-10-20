package service

import (
	"admin-permission/config"
	"admin-permission/model"
	"admin-permission/repository"
	"admin-permission/response"
	"errors"
	"fmt"
)

type LoginServicer interface {
	Login() (info response.LoginSuccess, err error)
}

type PasswordLogin struct {
	Account   string
	Password  string
	adminRepo *repository.AdminRepository
}

type CodeLogin struct {
	Account   string
	Code      string
	adminRepo *repository.AdminRepository
}

func NewLogin(account, password string) map[string]LoginServicer {
	return map[string]LoginServicer{
		"password": NewPasswordLogin(account, password),
		"code":     NewCodeLogin(account, password),
	}
}

func NewPasswordLogin(account, password string) *PasswordLogin {
	return &PasswordLogin{
		Account:   account,
		Password:  password,
		adminRepo: repository.NewAdminRepository(),
	}
}

func NewCodeLogin(account, password string) *CodeLogin {
	return &CodeLogin{
		Account:   account,
		Code:      password,
		adminRepo: repository.NewAdminRepository(),
	}
}

func getUserByAccount(account string, repo *repository.AdminRepository) (user model.AdminModel, err error) {
	user, err = repo.GetByAccount(account, config.ADMIN_STATUS_FRZEEZE, config.ADMIN_STATUS_NORMAL)
	return
}

func checkUserStatus(user model.AdminModel) 

func (pl *PasswordLogin) Login() (info response.LoginSuccess, err error) {
	user, err := pl.adminRepo.GetByAccount(pl.Account, config.ADMIN_STATUS_FRZEEZE, config.ADMIN_STATUS_NORMAL)
	if err != nil {
		return info, err
	}
	
	return
}

func (cl *CodeLogin) Login() (info response.LoginSuccess, err error) {
	user, err := cl.adminRepo.GetByAccount(cl.Account, config.ADMIN_STATUS_FRZEEZE, config.ADMIN_STATUS_NORMAL)
	if err != nil {

	}
	fmt.Println(user)
	return
}
