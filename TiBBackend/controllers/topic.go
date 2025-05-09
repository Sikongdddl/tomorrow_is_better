package controllers

import (
	"gorm.io/gorm"
	"net/http"
	"time"

	"TiBBackend/config"
	"TiBBackend/models"
	"github.com/gin-gonic/gin"
)

type CreateTopicInput struct {
	Title     string    `json:"title" binding:"required"`
	EventTime time.Time `json:"event_time" binding:"required"`
	Location  string    `json:"location"`
	CreatorID uint      `json:"creator_id" binding:"required"`
}

func CreateTopic(c *gin.Context) {
	var input CreateTopicInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	topic := models.Topic{
		Title:     input.Title,
		EventTime: input.EventTime,
		Location:  input.Location,
		CreatorID: input.CreatorID,
	}

	if err := config.DB.Create(&topic).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 自动添加发起者为参与者
	participant := models.TopicParticipant{
		TopicID:  topic.ID,
		UserID:   input.CreatorID,
		Role:     "creator",
		JoinedAt: time.Now(),
	}
	config.DB.Create(&participant)

	c.JSON(http.StatusOK, gin.H{"message": "topic created", "topic": topic})
}

func JoinTopic(c *gin.Context) {
	var input struct {
		UserID  uint `json:"user_id" binding:"required"`
		TopicID uint `json:"topic_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查是否已加入
	var existing models.TopicParticipant
	result := config.DB.Where("topic_id = ? AND user_id = ?", input.TopicID, input.UserID).First(&existing)
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already joined this topic"})
		return
	}

	participant := models.TopicParticipant{
		TopicID:  uint(input.TopicID),
		UserID:   input.UserID,
		Role:     "participant",
		JoinedAt: time.Now(),
	}
	if err := config.DB.Create(&participant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Joined topic"})
}

func LeaveTopic(c *gin.Context) {
	type LeaveTopicRequest struct {
		UserID  uint `json:"user_id" binding:"required"`
		TopicID uint `json:"topic_id" binding:"required"`
	}

	var req LeaveTopicRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// 查找参与者记录
	var participant models.TopicParticipant
	if err := config.DB.Where("user_id = ? AND topic_id = ?", req.UserID, req.TopicID).First(&participant).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User is not a participant of the topic"})
		return
	}

	// 删除参与者记录
	if err := config.DB.Delete(&participant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove participant from topic"})
		return
	}

	// 检查是否还有其他参与者
	var remainingCount int64
	if err := config.DB.Model(&models.TopicParticipant{}).
		Where("topic_id = ?", req.TopicID).
		Count(&remainingCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check remaining participants"})
		return
	}

	// 如果没有参与者，删除该活动
	if remainingCount == 0 {
		if err := config.DB.Delete(&models.Topic{}, req.TopicID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete empty topic"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully left topic"})
}

func ListTopics(c *gin.Context) {
	var topics []models.Topic
	if err := config.DB.Preload("Creator").Preload("Participants").Find(&topics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"topics": topics})
}

func GetTopicById(c *gin.Context) {
	var request struct {
		TopicID int `json:"topic_id"` // 从请求体获取 topic_id
	}

	// 绑定请求体中的 JSON 数据
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// 根据 topic_id 查询数据库
	var topic models.Topic
	if err := config.DB.First(&topic, request.TopicID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Topic not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": topic})
}
