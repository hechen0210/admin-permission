package model

import (
	"admin-permission/config"

	"gorm.io/gorm"
)

type MenuModel struct {
	base
	Title   string `gorm:"column:title;type:string;size:15;comment:菜单名称"`
	Parent  int    `gorm:"column:parent;type:int;size:5;comment:父菜单ID"`
	Level   int    `gorm:"column:level;type:int;size:1;comment:菜单等级"`
	Sort    int    `gorm:"column:sort;type:int;size:5;comment:排序"`
	Status  int    `gorm:"column:status;type:int;size:1;comment:状态"`
}

func (MenuModel) TableName() string {
	return config.GetDbPrefix() + "menu"
}

func NewMenuModel() *MenuModel {
	return &MenuModel{}
}

// 获取所有菜单
func (mm *MenuModel) GetAll() (list []MenuModel, err error) {
	err = config.GetDb().Find(&list).Error
	return list, err
}

// 通过菜单ID获取菜单信息
// id 菜单ID
func (mm *MenuModel) GetById(id int) (menu MenuModel, err error) {
	err = config.GetDb().First(&menu, id).Error
	return menu, err
}

// 通过菜单标识获取菜单信息
// mark 菜单标识
func (mm *MenuModel) GetByMark(mark string) (menu MenuModel, err error) {
	err = config.GetDb().Where("mark = ?", mark).First(&menu).Error
	return menu, err
}

func (mm *MenuModel) GetCount(condition map[string]interface{}) (count int64,err error) {
	err = mm.byWhere(condition).Model(mm).First(&count).Error
	return
}

// 获取菜单列表
func (mm *MenuModel) GetList(page,pageSize int,condition map[string]interface{}) (list []MenuModel, err error) {
	query := mm.byWhere(condition)
	query = pagintion(query, page, pageSize)
	err = query.Find(&list).Error
	return list, err
}

func (mm *MenuModel) byWhere(condition map[string]interface{}) (db *gorm.DB){
	db = config.GetDb()
    if title, ok := condition["title"]; ok {
        db = db.Scopes(mm.ByTitle(title))
    }
	return db
}

func (mm *MenuModel) ByTitle(title interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("title like (?)", title)
	}
}
