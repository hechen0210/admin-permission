package config

import "gorm.io/gorm"

type Config struct {
	Mysql    *gorm.DB
	DbPrefix string
}

var ConfigData Config

func New(config Config) {
	ConfigData = config
}

func GetDb() *gorm.DB {
	return ConfigData.Mysql
}

func GetDbPrefix() string {
	return ConfigData.DbPrefix
}
