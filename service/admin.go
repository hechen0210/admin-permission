package service

import (
	"admin-permission/config"
	"admin-permission/model"
	"admin-permission/repository"
	"fmt"
)

type AdminService struct {
	adminRepo *repository.AdminRepository
}

func NewAdminService() *AdminService {
	return &AdminService{
		adminRepo: repository.NewAdminRepository(),
	}
}

// 登录
func (s *AdminService) Login(account, password string) {
	user, err := s.adminRepo.GetByAccount(account, config.ADMIN_STATUS_NORMAL)
	if err != nil {

	}
	fmt.Println(user)
}

// 获取用户信息
func (s *AdminService) GetInfo(where map[string]interface{}) {
	var err error
	user := model.AdminModel{}
	if id, ok := where["id"]; ok {
		user, err = s.adminRepo.GetById(id.(int))
	}
	if account, ok := where["account"]; ok {
		user, err = s.adminRepo.GetByAccount(account.(string), config.ADMIN_STATUS_IGNORE)
	}
	fmt.Println(user, err)
}

// 获取用户总数
func (s *AdminService) GetCount(where map[string]interface{}) (count int64, err error) {
	return s.adminRepo.GetCount(where)
}

// 获取用户列表
func (s *AdminService) GetList(page, pageSize int, where map[string]interface{}) {

}

func (s *AdminService) Add() {

}

func (s *AdminService) Update() {

}

// 冻结账号
func (s *AdminService) Freeze(id int) error {
	return s.adminRepo.Update(map[string]interface{}{"id": id}, map[string]interface{}{"status": 0})
}

// 解冻账号
func (s *AdminService) UnFreeze(id int) error {
	return s.adminRepo.Update(map[string]interface{}{"id": id}, map[string]interface{}{"status": 1})
}

// 删除账号
func (s *AdminService) Delete(id int) error {
	return s.adminRepo.Update(map[string]interface{}{"id": id}, map[string]interface{}{"status": -1})
}
