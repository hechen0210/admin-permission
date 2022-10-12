package repository

import (
	"admin-permission/config"
	"admin-permission/model"
	"errors"

	"gorm.io/gorm"
)

type RouteRepository struct {
	db         *gorm.DB
}

func NewRouteRepository() *RouteRepository {
	return &RouteRepository{
		db:         config.GetDb(),
	}
}

func (rr *RouteRepository) GetAll() (list []model.RouteModel, err error) {
	err = rr.db.Find(&list).Error
	return
}

func (rr *RouteRepository) Add(data []string) {
	for _, item := range data {
		var route model.RouteModel
		err := rr.db.Where("route=?", item).Find(&route).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			rr.db.Create(&model.RouteModel{
				Url: item,
			})
		}
	}
}
