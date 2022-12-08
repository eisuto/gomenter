package main

import (
	"gomenter/models"
	"gomenter/routers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// 创建数据库
	db, err := gorm.Open(sqlite.Open("comments.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移数据库结构
	db.AutoMigrate(&models.Comment{})

	// 开启
	r := routers.InitRouter(db)
	r.Run(":8080")

}
