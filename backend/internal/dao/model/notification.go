package model

import "time"

type Notification struct {
	ID         int64     `gorm:"primaryKey;column:id"`
	UserID     int64     `gorm:"column:user_id;not null;index:notifications_user_id_idx"`
	Type       string    `gorm:"column:type;size:30;default:comment"`
	Title      string    `gorm:"column:title;size:200"`
	Content    string    `gorm:"column:content"`
	FromUserID *int64    `gorm:"column:from_user_id"`
	PostID     *int64    `gorm:"column:post_id"`
	IsRead     bool      `gorm:"column:is_read;default:false"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (Notification) TableName() string {
	return "notifications"
}
