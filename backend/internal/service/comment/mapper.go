package comment

import (
	"helloblog/internal/dao/model"
	"helloblog/internal/dto"
)

func toModel(userID, postID int64, req dto.CreateCommentRequest) *model.Comment {
	return &model.Comment{
		PostID:   postID,
		UserID:   userID,
		ParentID: req.ParentID,
		Content:  req.Content,
	}
}

func toResponse(c *model.Comment) dto.CommentResponse {
	resp := dto.CommentResponse{
		ID:        c.ID,
		PostID:    c.PostID,
		ParentID:  c.ParentID,
		Content:   c.Content,
		LikeCount: c.LikeCount,
		CreatedAt: c.CreatedAt,
	}

	if c.User != nil {
		resp.User = dto.UserResponse{
			ID:        c.User.ID,
			Username:  c.User.Username,
			AvatarURL: c.User.AvatarURL,
		}
	}

	if len(c.Children) > 0 {
		resp.Children = make([]dto.CommentResponse, len(c.Children))
		for i, child := range c.Children {
			resp.Children[i] = toResponse(child)
		}
	}

	return resp
}
