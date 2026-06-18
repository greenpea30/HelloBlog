package dto

import "time"

type CreateFolderRequest struct {
	Name string `json:"name" binding:"required,min=1,max=50"`
}

type FolderResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	PostCount int       `json:"post_count"`
	CreatedAt time.Time `json:"created_at"`
}

type UserProfileResponse struct {
	User    UserResponse      `json:"user"`
	Folders []FolderWithPosts `json:"folders"`
}

type FolderWithPosts struct {
	ID    int64          `json:"id"`
	Name  string         `json:"name"`
	Posts []PostResponse `json:"posts"`
}
