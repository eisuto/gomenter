package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Comment struct {
	gorm.Model
	Content string `json:"content"`
	Author  string `json:"author"`
}

// CreateComment 定义添加评论的处理函数
func CreateComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// 解析请求参数
	var comment Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 创建评论
	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{"message": "Comment created successfully"})
}

// DeleteComment 定义删除评论的处理函数
func DeleteComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// 解析请求参数
	var comment Comment
	if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// 删除评论
	if err := db.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

// ListComment 定义查看评论列表的处理函数
func ListComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// 查询所有评论
	var comments []Comment
	if err := db.Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{"data": comments})
}
