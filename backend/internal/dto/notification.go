package dto

import "time"

type NotificationResponse struct {
	ID         int64     `json:"id"`
	Type       string    `json:"type"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	FromUserID *int64    `json:"from_user_id,omitempty"`
	PostID     *int64    `json:"post_id,omitempty"`
	IsRead     bool      `json:"is_read"`
	CreatedAt  time.Time `json:"created_at"`
}
