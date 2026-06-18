package model

import "time"

type Folder struct {
	ID        int64     `gorm:"primaryKey;column:id"`
	UserID    int64     `gorm:"column:user_id;not null;index"`
	Name      string    `gorm:"column:name;size:50;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Folder) TableName() string {
	return "folders"
}
