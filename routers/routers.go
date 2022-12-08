package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gomenter/models"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// 注入db引用

	// 定义路由组
	commentGroup := r.Group("/comments")
	// 注册路由
	commentGroup.POST("/", injectDB(db), models.CreateComment)
	commentGroup.GET("/", injectDB(db), models.ListComment)
	commentGroup.DELETE("/:id", injectDB(db), models.DeleteComment)

	return r
}

func injectDB(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 将db引用注入到上下文中
		c.Set("db", db)
		c.Next()
	}
}
