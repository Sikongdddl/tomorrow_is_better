package config

import (
	"TiBBackend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "jrgong:123456@tcp(localhost:3306)/tib?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	DB = db

	// Auto migrate schema
	err = db.AutoMigrate(&models.User{}, &models.Topic{}, &models.TopicParticipant{}, &models.Comment{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
}
