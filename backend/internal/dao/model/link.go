package model

import "time"

type Link struct {
	ID        int64     `gorm:"primaryKey;column:id"`
	Name      string    `gorm:"column:name;size:100;not null"`
	URL       string    `gorm:"column:url;size:500;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (Link) TableName() string {
	return "links"
}
