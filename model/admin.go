package model

import (
	"admin-permission/config"

	"gorm.io/gorm"
)

type AdminModel struct {
	base
	Account       string `gorm:"column:account;type:string;size:20;index;unique;not null;comment:账号"`
	Password      string `gorm:"column:password;type:string;size:64;not null;comment:密码"`
	Name          string `gorm:"column:name;type:string;size:10;comment:真实姓名"`
	Mobile        string `gorm:"column:mobile;type:string;size:11;comment:手机号码"`
	Email         string `gorm:"column:email;type:string;size:30;comment:电子邮箱"`
	Group         int    `gorm:"column:group;type:int;size:5;default:0;index;comment:用户组"`
	Role          int    `gorm:"column:role;type:int;size:5;default:0"`
	Status        int    `gorm:"column:status;type:tinyInt;size:5;default:1;comment:状态，0-冻结，1-可用，-1-删除"`
	LastLoginIp   int    `gorm:"column:last_login_ip;type:int;size:10;comment:最后登录IP"`
	LastLoginTime int    `gorm:"column:last_login_time;type:int;size:10;comment:最后登录时间"`
}

func (AdminModel) TableName() string {
	return config.GetDbPrefix() + "admin"
}

func NewAdminModel(base AdminModel) *AdminModel {
	return &AdminModel{}
}

func (am *AdminModel) GetById(id int) (admin AdminModel, err error) {
	err = config.GetDb().First(&admin, id).Error
	return
}

func (am *AdminModel) GetByAccount(account string) (admin AdminModel, err error) {
	err = config.GetDb().Where("account = ?", account).First(&admin).Error
	return
}

func (am *AdminModel) GetCount(condition map[string]interface{}) (count int64, err error) {
	err = am.byWhere(condition).Model(am).Count(&count).Error
	return
}

func (am *AdminModel) GetList(page, pageSize int, condition map[string]interface{}) (list []AdminModel, err error) {
	query := am.byWhere(condition)
	query = pagintion(query, page, pageSize)
	err = query.Find(&list).Error
	return
}

func (am *AdminModel) byWhere(condition map[string]interface{}) (db *gorm.DB) {
	db = config.GetDb()
	if account, ok := condition["account"]; ok {
		db = db.Scopes(am.ByAccount(account.(string)))
	}
	if name, ok := condition["name"]; ok {
		db = db.Scopes(am.ByName(name.(string)))
	}
	if group, ok := condition["group"]; ok {
		db = db.Scopes(am.ByGroup(group.(int)))
	}
	return db
}

func (am *AdminModel) ByAccount(account string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("account like ?", "%"+account+"%")
	}
}

func (am *AdminModel) ByName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name like ?", "%"+name+"%")
	}
}

func (am *AdminModel) ByGroup(group int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("group = ?", group)
	}
}
