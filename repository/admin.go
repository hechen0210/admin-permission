package repository

import (
	"admin-permission/config"
	"admin-permission/model"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{
		db: config.GetDb(),
	}
}

func (ar *AdminRepository) GetById(id int) (admin model.AdminModel, err error) {
	err = ar.db.First(&admin, id).Error
	return
}

func (ar *AdminRepository) GetByAccount(account string, status ...int) (admin model.AdminModel, err error) {
	query := ar.db.Where("account = ?", account)
	if len(status) > 0 && status[0] != config.ADMIN_STATUS_IGNORE {
		query = query.Where("status in (?)", status)
	}
	err = query.First(&admin).Error
	return
}

func (ar *AdminRepository) GetCount(condition map[string]interface{}) (count int64, err error) {
	err = ar.byWhere(condition).Model(&model.AdminModel{}).Count(&count).Error
	return
}

func (ar *AdminRepository) GetList(page, pageSize int, condition map[string]interface{}) (list []model.AdminModel, err error) {
	query := ar.byWhere(condition)
	query = pagintion(query, page, pageSize)
	err = query.Find(&list).Error
	return
}

func (ar *AdminRepository) Update(where, data map[string]interface{}) error {
	return ar.db.Model(ar.adminModel).Where(where).Updates(data).Error
}

func (ar *AdminRepository) byWhere(condition map[string]interface{}) (db *gorm.DB) {
	db = ar.db
	if account, ok := condition["account"]; ok {
		db = db.Scopes(ar.ByAccount(account.(string)))
	}
	if name, ok := condition["name"]; ok {
		db = db.Scopes(ar.ByName(name.(string)))
	}
	if group, ok := condition["group"]; ok {
		db = db.Scopes(ar.ByGroup(group.(int)))
	}
	return db
}

func (ar *AdminRepository) ByAccount(account string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("account like ?", "%"+account+"%")
	}
}

func (ar *AdminRepository) ByName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name like ?", "%"+name+"%")
	}
}

func (ar *AdminRepository) ByGroup(group int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("group = ?", group)
	}
}
