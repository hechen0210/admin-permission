package repository

import (
	"admin-permission/config"
	"admin-permission/model"

	"gorm.io/gorm"
)

type AdminPrivilegeRepository struct {
	db *gorm.DB
}

func NewAdminPrivilegeRepository() *AdminPrivilegeRepository {
	return &AdminPrivilegeRepository{
		db: config.GetDb(),
	}
}

func (apr *AdminPrivilegeRepository) GetById(id int) (info model.AdminPrivilegeModel, err error) {
	err = config.GetDb().Find(&info, id).Error
	return
}

func (apr *AdminPrivilegeRepository) GetByAdminId(adminId int) (info model.AdminPrivilegeModel, err error) {
	err = config.GetDb().Where("admin_id = ?", adminId).Find(&info).Error
	return
}
