package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(host string, name string) {
	var err error

	// 创建数据库
	db, err = gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移数据库结构
	db.AutoMigrate(&Comment{})

}
