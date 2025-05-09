package models

import "time"

type User struct {
	ID                      uint   `gorm:"primaryKey"`
	Username                string `gorm:"type:varchar(64);uniqueIndex;not null"`
	PasswordHash            string `gorm:"type:varchar(255);not null"`
	AvatarURL               string
	ViolationCount          int
	SuccessParticipationCnt int
	CreatedAt               time.Time
}

type Topic struct {
	ID           uint      `gorm:"primaryKey"`
	Title        string    `gorm:"not null"`
	EventTime    time.Time `gorm:"not null"`
	Location     string
	CreatorID    uint
	Creator      User `gorm:"foreignKey:CreatorID"`
	CreatedAt    time.Time
	Participants []TopicParticipant `gorm:"foreignKey:TopicID"`
}

type TopicParticipant struct {
	ID       uint `gorm:"primaryKey"`
	TopicID  uint
	UserID   uint
	Role     string `gorm:"type:enum('creator','participant','past_participant')"`
	JoinedAt time.Time
}

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	TopicID   uint   `gorm:"not null"`
	Content   string `gorm:"type:text;not null"`
	CreatedAt time.Time
	
	User  User  `gorm:"foreignKey:UserID"`
	Topic Topic `gorm:"foreignKey:TopicID"`
}
