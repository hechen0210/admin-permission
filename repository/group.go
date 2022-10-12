package repository

import (
	"admin-permission/config"
	"admin-permission/model"

	"gorm.io/gorm"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository() *GroupRepository {
	return &GroupRepository{
		db: config.GetDb(),
	}
}


func (gr *GroupRepository) GetById(id int) (info model.GroupModel, err error) {
	err = config.GetDb().First(&info, id).Error
	return
}

func (gr *GroupRepository) GetCount(condition map[string]interface{}) (count int64, err error) {
	err = gr.byWhere(condition).Model(&model.GroupModel{}).Count(&count).Error
	return
}

func (gr *GroupRepository) GetList(page, pageSize int, condition map[string]interface{}) (list []model.GroupModel, err error) {
	query := gr.byWhere(condition)
	query = pagintion(query, page, pageSize)
	err = query.Find(&list).Error
	return
}

func (gr *GroupRepository) byWhere(condition map[string]interface{}) (db *gorm.DB) {
	db = config.GetDb()
	if name, ok := condition["name"]; ok {
		db = db.Scopes(gr.ByName(name.(string)))
	}
	return db
}

func (gr *GroupRepository) ByName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name like ?", "%"+name+"%")
	}
}


