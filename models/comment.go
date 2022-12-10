package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Comment struct {
	gorm.Model
	ArticleId string `json:"articleId"`
	Content   string `json:"content"`
	Author    string `json:"author"`
}

// CreateComment 保存
func CreateComment(c *gin.Context) {
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

// DeleteComment 删除
func DeleteComment(c *gin.Context) {
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

// ListComment 列表
func ListComment(c *gin.Context) {
	// 根据articleId或author查询评论
	var comments []Comment
	if err := db.Where("article_id = ?", c.Query("articleId")).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 返回响应
	c.JSON(http.StatusOK, gin.H{"data": comments})
}
