package repository

import (
	"admin-permission/config"
	"admin-permission/model"

	"gorm.io/gorm"
)

type MenuRepository struct {
	db *gorm.DB
}

func NewMenuRepository() *MenuRepository {
	return &MenuRepository{
		db: config.GetDb(),
	}
}


// 获取所有菜单
func (mr *MenuRepository) GetAll() (list []model.MenuModel, err error) {
	err = mr.db.Find(&list).Error
	return list, err
}

// 通过菜单ID获取菜单信息
// id 菜单ID
func (mr *MenuRepository) GetById(id int) (menu model.MenuModel, err error) {
	err = mr.db.First(&menu, id).Error
	return menu, err
}

// 通过菜单标识获取菜单信息
// mark 菜单标识
func (mr *MenuRepository) GetByMark(mark string) (menu model.MenuModel, err error) {
	err = mr.db.Where("mark = ?", mark).First(&menu).Error
	return menu, err
}

func (mr *MenuRepository) GetCount(condition map[string]interface{}) (count int64, err error) {
	err = mr.byWhere(condition).Model(&model.MenuModel{}).First(&count).Error
	return
}

// 获取菜单列表
func (mr *MenuRepository) GetList(page, pageSize int, condition map[string]interface{}) (list []model.MenuModel, err error) {
	query := mr.byWhere(condition)
	query = pagintion(query, page, pageSize)
	err = query.Find(&list).Error
	return list, err
}

func (mr *MenuRepository) byWhere(condition map[string]interface{}) (db *gorm.DB) {
	db = mr.db
	if title, ok := condition["title"]; ok {
		db = db.Scopes(mr.ByTitle(title))
	}
	return db
}

func (mr *MenuRepository) ByTitle(title interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("title like (?)", title)
	}
}
