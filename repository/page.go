package repository

import (
	"admin-permission/config"
	"admin-permission/model"

	"gorm.io/gorm"
)

type PageRepository struct {
	db *gorm.DB
}

func NewPageRepository() *PageRepository {
	return &PageRepository{
		db: config.GetDb(),
	}
}


func (pp *PageRepository) GetById(id int) (info model.PageModel, err error) {
	err = pp.db.Where("id =?", id).First(&info).Error
	return
}

func (pp *PageRepository) GetCount(condition map[string]interface{}) (count int64, err error) {
	err = pp.byWhere(condition).Model(&model.PageModel{}).Count(&count).Error
	return
}

func (pp *PageRepository) GetList(page, pageSize int, condition map[string]interface{}) (list []model.PageModel, err error) {
	query := pp.byWhere(condition)
	query = pagintion(query, page, pageSize)
	err = query.Find(&list).Error
	return
}

func (pp *PageRepository) byWhere(condition map[string]interface{}) (db *gorm.DB) {
	db = pp.db
	if name, ok := condition["name"]; ok {
		db = db.Scopes(pp.byName(name.(string)))
	}
	if pageType, ok := condition["type"]; ok {
		db = db.Scopes(pp.byType(pageType.(int)))
	}
	if parent, ok := condition["parent"]; ok {
		db = db.Scopes(pp.byParent(parent.(int)))
	}
	return
}

func (pp *PageRepository) byName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name like ?", "%"+name+"%")
	}
}

func (pp *PageRepository) byType(pageType int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("type = ?", pageType)
	}
}

func (pp *PageRepository) byParent(parent int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("parent = ? ", parent)
	}
}
