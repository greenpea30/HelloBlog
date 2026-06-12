package dto

import "time"

type CreateLinkRequest struct {
	Name string `json:"name" binding:"required,min=1,max=100"`
	URL  string `json:"url" binding:"required,min=1,max=500"`
}

type LinkResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
