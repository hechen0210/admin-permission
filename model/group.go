package model

import (
	"admin-permission/config"

	"gorm.io/gorm"
)

type GroupModel struct {
	base
	Name          string `gorm:"column:name;type:string;size:20;not null;comment:权限组名称"`
	Mark          string `gorm:"column:mark;type:string;size:20;not null;comment:权限组标识，唯一"`
	Parent        int    `gorm:"column:parent;type:int;size:5;default:0;not null;comment:上级权限组"`
	DataPrivilege int    `gorm:"column:data_privilege;type:tinyInt;size:1;default:0;comment:数据权限，0-所有数据，1-该组数据，2-仅自己的数据"`
	Privilege     string `gorm:"column:privilege;type:text;comment:权限表"`
}

func (GroupModel) TableName() string {
	return config.GetDbPrefix() + "groups"
}

func NewGroupModel() *GroupModel {
	return &GroupModel{}
}

func (gm *GroupModel) GetById(id int) (info GroupModel, err error) {
	err = config.GetDb().First(&info, id).Error
	return
}

func (gm *GroupModel) GetCount(condition map[string]interface{}) (count int64, err error) {
	err = gm.byWhere(condition).Model(gm).Count(&count).Error
	return
}

func (gm *GroupModel) GetList(page, pageSize int, condition map[string]interface{}) (list []GroupModel, err error) {
	query := gm.byWhere(condition)
	query = pagintion(query, page, pageSize)
	err = query.Find(&list).Error
	return
}

func (gm *GroupModel) byWhere(condition map[string]interface{}) (db *gorm.DB) {
	db = config.GetDb()
	if name, ok := condition["name"]; ok {
		db = db.Scopes(gm.ByName(name.(string)))
	}
	return db
}

func (gm *GroupModel) ByName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name like ?", "%"+name+"%")
	}
}
