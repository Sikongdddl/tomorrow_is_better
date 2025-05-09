package controllers

import (
	"TiBBackend/config"
	"TiBBackend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AddComment(c *gin.Context) {
	type AddCommentRequest struct {
		UserID  uint   `json:"user_id" binding:"required"`
		TopicID uint   `json:"topic_id" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	var req AddCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	comment := models.Comment{
		UserID:    req.UserID,
		TopicID:   req.TopicID,
		Content:   req.Content,
		CreatedAt: time.Now(),
	}

	if err := config.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment added successfully", "comment_id": comment.ID})
}

func DeleteComment(c *gin.Context) {
	type DeleteCommentRequest struct {
		CommentID uint `json:"comment_id" binding:"required"`
	}

	var req DeleteCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := config.DB.Delete(&models.Comment{}, req.CommentID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

func GetCommentsByTopic(c *gin.Context) {
	type GetCommentsRequest struct {
		TopicID uint `json:"topic_id" binding:"required"`
	}

	var req GetCommentsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var comments []models.Comment
	if err := config.DB.
		Where("topic_id = ?", req.TopicID).
		Preload("User").
		Order("created_at ASC").
		Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	// 返回用户友好的格式（用户名+内容+时间）
	var result []gin.H
	for _, cmt := range comments {
		result = append(result, gin.H{
			"comment_id": cmt.ID,
			"user_id":    cmt.UserID,
			"username":   cmt.User.Username,
			"content":    cmt.Content,
			"created_at": cmt.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"comments": result,
	})
}
