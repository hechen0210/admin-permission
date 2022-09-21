package model

import "admin-permission/config"

type RouteModel struct {
	base
	Url string `gorm:"column:url,type:string,size:64"`
}

func (RouteModel) TableName() string {
	return config.GetDbPrefix() + "route"
}

func NewRouteModel() *RouteModel {
	return &RouteModel{}
}

func (m *RouteModel) GetAll() (list []RouteModel, err error) {
	err = config.GetDb().Find(&list).Error
	return
}
