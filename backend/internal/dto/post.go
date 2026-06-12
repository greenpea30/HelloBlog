package dto

import "time"

// 请求
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=200"`
	Summary string `json:"summary" binding:"max=500"`
	Content string `json:"content" binding:"required,min=1"`
}

type UpdatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=200"`
	Summary string `json:"summary" binding:"max=500"`
	Content string `json:"content" binding:"required,min=1"`
}

type PostListRequest struct {
	Page     int    `form:"page" binding:"omitempty,min=1"`
	PageSize int    `form:"page_size" binding:"omitempty,min=1,max=50"`
	UserID   int64  `form:"user_id"`
	OrderBy  string `form:"order_by"`
	ZJUOnly  bool   `form:"zju_only"`
}

// 响应
type PostResponse struct {
	ID           int64        `json:"id"`
	Title        string       `json:"title"`
	Summary      string       `json:"summary"`
	Content      string       `json:"content,omitempty"`
	User         UserResponse `json:"user,omitempty"`
	LikeCount    int          `json:"like_count"`
	CommentCount int          `json:"comment_count"`
	ViewCount    int          `json:"view_count"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type PostListResponse struct {
	Items      []PostResponse `json:"items"`
	Total      int64          `json:"total"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	TotalPages int64          `json:"total_pages"`
}
