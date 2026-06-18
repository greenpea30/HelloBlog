package post

import (
	"helloblog/internal/dao/model"
	"helloblog/internal/dto"
)

func toModel(userID int64, req dto.CreatePostRequest) *model.Post {
	format := req.Format
	if format == "" {
		format = "markdown"
	}
	return &model.Post{
		UserID:  userID,
		Title:   req.Title,
		Summary: req.Summary,
		Content: req.Content,
		Format:  format,
	}
}

func toResponse(post *model.Post) dto.PostResponse {
	format := post.Format
	if format == "" {
		format = "markdown"
	}
	resp := dto.PostResponse{
		ID:           post.ID,
		Title:        post.Title,
		Summary:      post.Summary,
		Content:      post.Content,
		Format:       format,
		LikeCount:    post.LikeCount,
		CommentCount: post.CommentCount,
		ViewCount:    post.ViewCount,
		CreatedAt:    post.CreatedAt,
		UpdatedAt:    post.UpdatedAt,
	}

	if post.User != nil {
		resp.User = dto.UserResponse{
			ID:        post.User.ID,
			Username:  post.User.Username,
			AvatarURL: post.User.AvatarURL,
		}
	}

	return resp
}
