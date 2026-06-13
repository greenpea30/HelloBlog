package dto

import "time"

// 请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"omitempty,email,max=100"`
	Password string `json:"password" binding:"required,min=8,max=72"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"omitempty,email,max=100"`
	Password string `json:"password" binding:"required,min=8,max=72"`
}

type ZJULoginRequest struct {
	StudentID string `json:"student_id" binding:"required,min=5,max=20"`
	Password  string `json:"password" binding:"required,min=1,max=72"`
}

type UpdateProfileRequest struct {
	Username  string `json:"username" binding:"required,min=3,max=50"`
	AvatarURL string `json:"avatar_url" binding:"omitempty,max=500"`
	Bio       string `json:"bio" binding:"omitempty,max=200"`
}

// 响应
type UserResponse struct {
	ID        int64     `json:"id"`
	ZjuID     *string   `json:"zju_id,omitempty"`
	Username  string    `json:"username"`
	Email     *string   `json:"email,omitempty"`
	AvatarURL string    `json:"avatar_url,omitempty"`
	Bio       string    `json:"bio,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type AuthResponse struct {
	User        UserResponse `json:"user"`
	AccessToken string       `json:"access_token"`
}
