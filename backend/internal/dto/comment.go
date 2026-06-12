package dto

import "time"

type CreateCommentRequest struct {
	Content  string `json:"content" binding:"required,min=1,max=2000"`
	ParentID *int64 `json:"parent_id"`
}

type CommentResponse struct {
	ID        int64             `json:"id"`
	PostID    int64             `json:"post_id"`
	User      UserResponse      `json:"user"`
	ParentID  *int64            `json:"parent_id,omitempty"`
	Content   string            `json:"content"`
	LikeCount int               `json:"like_count"`
	Children  []CommentResponse `json:"children,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
}
