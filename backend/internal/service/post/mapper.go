package post

import (
	"helloblog/internal/dao/model"
	"helloblog/internal/dto"
)

func toModel(userID int64, req dto.CreatePostRequest) *model.Post {
	return &model.Post{
		UserID:  userID,
		Title:   req.Title,
		Summary: req.Summary,
		Content: req.Content,
	}
}

func toResponse(post *model.Post) dto.PostResponse {
	resp := dto.PostResponse{
		ID:           post.ID,
		Title:        post.Title,
		Summary:      post.Summary,
		Content:      post.Content,
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
