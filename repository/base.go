package repository

import "gorm.io/gorm"

func pagintion(db *gorm.DB, page, pageSize int) *gorm.DB {
	if pageSize > 0 {
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	return db
}
