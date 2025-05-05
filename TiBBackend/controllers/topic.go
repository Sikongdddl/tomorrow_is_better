package controllers

import (
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

func ListTopics(c *gin.Context) {
	var topics []models.Topic
	if err := config.DB.Preload("Creator").Preload("Participants").Find(&topics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"topics": topics})
}
