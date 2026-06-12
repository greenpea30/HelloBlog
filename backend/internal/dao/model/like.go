package model

import "time"

type Like struct {
	ID         int64     `gorm:"primaryKey;column:id"`
	UserID     int64     `gorm:"column:user_id;not null"`
	TargetType string    `gorm:"column:target_type;not null"`
	TargetID   int64     `gorm:"column:target_id;not null"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (Like) TableName() string {
	return "likes"
}
