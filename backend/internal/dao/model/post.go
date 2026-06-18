package model

import "time"

type Post struct {
	ID           int64     `gorm:"primaryKey;column:id"`
	UserID       int64     `gorm:"column:user_id;not null;index"`
	Title        string    `gorm:"column:title;size:200;not null"`
	Summary      string    `gorm:"column:summary"`
	Content      string    `gorm:"column:content;not null"`
	Format       string    `gorm:"column:format;size:10;default:markdown"`
	SearchVector string    `gorm:"column:search_vector;->"`
	LikeCount    int       `gorm:"column:like_count;default:0"`
	CommentCount int       `gorm:"column:comment_count;default:0"`
	ViewCount    int       `gorm:"column:view_count;default:0"`
	Status       string    `gorm:"column:status;default:normal"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`

	// 关联
	User *User `gorm:"foreignKey:UserID"`
}

func (Post) TableName() string {
	return "posts"
}

type PostEmbedding struct {
	ID        int64     `gorm:"primaryKey;column:id"`
	PostID    int64     `gorm:"column:post_id;unique;not null"`
	Embedding string    `gorm:"column:embedding;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (PostEmbedding) TableName() string {
	return "post_embeddings"
}
