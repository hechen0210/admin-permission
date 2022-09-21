package model

import "gorm.io/gorm"

type base struct {
	Id        int   `gorm:"column:id;primaryKey;autoIncrement;type:int"`
	CreatedAt int64 `gorm:"column:created_at;type:int;size:10;autoCreateTime"`
	UpdatedAt int64 `gorm:"column:updated_at;type:int;size:10;autoUpdateTime"`
}

func pagintion(db *gorm.DB, page, pageSize int) *gorm.DB {
	if pageSize > 0 {
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	return db
}
