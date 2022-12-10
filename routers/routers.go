package routers

import (
	"github.com/gin-gonic/gin"
	"gomenter/models"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 定义路由组
	commentGroup := r.Group("/comments")

	// 注册路由
	commentGroup.POST("/", models.CreateComment)
	commentGroup.GET("/", models.ListComment)
	commentGroup.DELETE("/:id", models.DeleteComment)

	return r
}
