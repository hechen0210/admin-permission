package model

import "admin-permission/config"

type AdminPrivilegeModel struct {
	base
	AdminId   int    `gorm:"column:admin_id;type:int;size:5;not null;comment:管理员ID"`
	Privilege string `gorm:"column:privilege;type:text;comment:权限表"`
}

func (AdminPrivilegeModel) TableName() string {
	return config.GetDbPrefix() + "admin_privileges"
}

func NewAdminPrivilegeModel() *AdminPrivilegeModel {
	return &AdminPrivilegeModel{}
}

func (apm *AdminPrivilegeModel) GetById(id int) (info AdminPrivilegeModel, err error) {
	err = config.GetDb().Find(&info, id).Error
	return
}

func (apm *AdminPrivilegeModel) GetByAdminId(adminId int) (info AdminPrivilegeModel, err error) {
	err = config.GetDb().Where("admin_id = ?", adminId).Find(&info).Error
	return
}
