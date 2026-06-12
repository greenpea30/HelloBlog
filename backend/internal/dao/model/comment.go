package model

import "time"

type Comment struct {
	ID        int64     `gorm:"primaryKey;column:id"`
	PostID    int64     `gorm:"column:post_id;not null;index"`
	UserID    int64     `gorm:"column:user_id;not null;index"`
	ParentID  *int64    `gorm:"column:parent_id;index"`
	Content   string    `gorm:"column:content;not null"`
	LikeCount int       `gorm:"column:like_count;default:0"`
	Status    string    `gorm:"column:status;default:normal"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`

	// 关联
	User     *User      `gorm:"foreignKey:UserID"`
	Post     *Post      `gorm:"foreignKey:PostID"`
	Parent   *Comment   `gorm:"foreignKey:ParentID"`
	Children []*Comment `gorm:"foreignKey:ParentID"`
}

func (Comment) TableName() string {
	return "comments"
}
