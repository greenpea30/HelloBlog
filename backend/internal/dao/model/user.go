package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           int64          `gorm:"primaryKey;column:id"`
	ZjuID        *string        `gorm:"column:zju_id;size:20;unique"`
	Username     string         `gorm:"column:username;size:50;unique;not null"`
	Email        *string        `gorm:"column:email;size:100;unique"`
	PasswordHash string         `gorm:"column:password_hash;size:255"`
	AvatarURL    string         `gorm:"column:avatar_url"`
	Bio          string         `gorm:"column:bio"`
	CreatedAt    time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (User) TableName() string {
	return "users"
}
