package model

import (
	"admin-permission/config"

	"gorm.io/gorm"
)

type PageModel struct {
	base
	Name    string `gorm:"column:name;type:string;size:20;not null;comment:页面名称"`
	Mark    string `gorm:"column:mark;type:string;size:20;unique;comment:页面标识，唯一"`
	PageUrl string `gorm:"column:page_url;type:string;size:64;comment:页面URL"`
	ApiUrl  string `gorm:"column:api_url;type:string;size:64;comment:api url"`
	Type    int    `gorm:"column:type;type:tinyInt;size:1;comment:页面类型，1-页面，2-功能"`
	Parent  int    `gorm:"column:parent;type:int;size:5;default:0;comment:父页面"`
}

func (PageModel) TableName() string {
	return config.GetDbPrefix() + "pages"
}

func NewPageModel() *PageModel {
	return &PageModel{}
}

func (pm *PageModel) GetById(id int) (info PageModel, err error) {
	err = config.GetDb().Where("id =?", id).First(&info).Error
	return
}

func (pm *PageModel) GetCount(condition map[string]interface{}) (count int64, err error) {
	err = pm.byWhere(condition).Model(pm).Count(&count).Error
	return
}

func (pm *PageModel) GetList(page, pageSize int, condition map[string]interface{}) (list []PageModel, err error) {
	query := pm.byWhere(condition)
	query = pagintion(query, page, pageSize)
	err = query.Find(&list).Error
	return
}

func (pm *PageModel) byWhere(condition map[string]interface{}) (db *gorm.DB) {
	db = config.GetDb()
	if name, ok := condition["name"]; ok {
		db = db.Scopes(pm.byName(name.(string)))
	}
	if pageType, ok := condition["type"]; ok {
		db = db.Scopes(pm.byType(pageType.(int)))
	}
	if parent, ok := condition["parent"]; ok {
		db = db.Scopes(pm.byParent(parent.(int)))
	}
	return
}

func (pm *PageModel) byName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name like ?", "%"+name+"%")
	}
}

func (pm *PageModel) byType(pageType int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("type = ?", pageType)
	}
}

func (pm *PageModel) byParent(parent int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("parent = ? ", parent)
	}
}
