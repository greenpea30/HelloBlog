package comment

import (
	"helloblog/internal/dto"
	"helloblog/internal/pkg/response"
)

func (s *Service) Create(userID, postID int64, req dto.CreateCommentRequest) (*dto.CommentResponse, error) {
	comment, err := s.comments.Create(toModel(userID, postID, req))
	if err != nil {
		return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
	}

	// 更新文章评论数
	if s.postCounter != nil {
		_ = s.postCounter.IncrementCommentCount(postID)
	}

	// 发送通知给文章作者
	if s.notifier != nil {
		s.notifier.NotifyComment(postID, userID, req.Content)
	}

	resp := toResponse(comment)
	return &resp, nil
}

func (s *Service) ListByPost(postID int64) ([]dto.CommentResponse, error) {
	comments, err := s.comments.ListByPost(postID, nil)
	if err != nil {
		return nil, response.Wrap(response.CodeInternalError, "internal server error", err)
	}

	resp := make([]dto.CommentResponse, len(comments))
	for i, c := range comments {
		resp[i] = toResponse(&c)
	}
	return resp, nil
}

func (s *Service) Delete(id, userID int64) error {
	return s.comments.SoftDelete(id, userID)
}
