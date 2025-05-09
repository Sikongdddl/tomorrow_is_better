package controllers

import (
	"TiBBackend/config"
	"TiBBackend/models"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := models.User{
		Username:     req.Username,
		PasswordHash: string(hash),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Username may already exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registered successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}

func LoginUser(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"avatar":   user.AvatarURL,
		},
	})
}

func GetNameByID(c *gin.Context) {
	var req struct {
		UserID uint `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, req.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": user.Username})
}

func GetUserInfoByID(c *gin.Context) {
	var req struct {
		UserID uint `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, req.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":                    user.ID,
		"username":              user.Username,
		"avatar_url":            user.AvatarURL,
		"violation_count":       user.ViolationCount,
		"success_participation": user.SuccessParticipationCnt,
		"created_at":            user.CreatedAt,
	})
}

func UploadAvatar(c *gin.Context) {
	// 解析 user_id
	userIDStr := c.PostForm("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	// 获取文件
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Avatar file required"})
		return
	}

	// 确保目录存在
	saveDir := "./uploads/avatars"
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create directory"})
		return
	}

	// 删除以 userID 开头的旧文件
	prefix := fmt.Sprintf("user_%d", userID)
	files, err := os.ReadDir(saveDir)
	if err == nil {
		for _, f := range files {
			if strings.HasPrefix(f.Name(), prefix) {
				os.Remove(filepath.Join(saveDir, f.Name()))
			}
		}
	}

	// 生成唯一新文件名（包含时间戳）
	timestamp := time.Now().Format("20060102150405") // 示例：20250506153045
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("user_%d_%s%s", userID, timestamp, ext)
	filePath := filepath.Join(saveDir, filename)

	// 保存文件
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save file"})
		return
	}

	// 构建访问路径
	avatarURL := "/static/avatars/" + filename

	// 更新数据库
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	user.AvatarURL = avatarURL
	config.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message":    "Avatar uploaded successfully",
		"avatar_url": avatarURL,
	})
}

func GetParticipatedTopics(c *gin.Context) {
	var req struct {
		UserID uint `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var participantEntries []models.TopicParticipant
	if err := config.DB.Where("user_id = ?", req.UserID).Find(&participantEntries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query participation"})
		return
	}

	var topicIDs []uint
	for _, entry := range participantEntries {
		topicIDs = append(topicIDs, entry.TopicID)
	}

	var topics []models.Topic
	if err := config.DB.Where("id IN ?", topicIDs).Find(&topics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get topics"})
		return
	}

	c.JSON(http.StatusOK, topics)
}

func GetCreatedTopics(c *gin.Context) {
	var req struct {
		UserID uint `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var topics []models.Topic
	if err := config.DB.Where("creator_id = ?", req.UserID).Find(&topics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get topics"})
		return
	}

	c.JSON(http.StatusOK, topics)
}
