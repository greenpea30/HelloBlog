package comment

import (
	"helloblog/internal/dao/model"
	"helloblog/internal/dto"
)

type UseCase interface {
	Create(userID, postID int64, req dto.CreateCommentRequest) (*dto.CommentResponse, error)
	ListByPost(postID int64) ([]dto.CommentResponse, error)
	Delete(id, userID int64) error
}

type commentRepository interface {
	Create(comment *model.Comment) (*model.Comment, error)
	ListByPost(postID int64, parentID *int64) ([]model.Comment, error)
	SoftDelete(id int64, userID int64) error
}

// PostCounter 用于更新文章评论数
type PostCounter interface {
	IncrementCommentCount(id int64) error
}

// Notifier 用于发送评论通知
type Notifier interface {
	NotifyComment(postID, fromUserID int64, content string)
}

type Service struct {
	comments    commentRepository
	postCounter PostCounter
	notifier    Notifier
}

func NewService(comments commentRepository, postCounter PostCounter, notifier Notifier) *Service {
	return &Service{comments: comments, postCounter: postCounter, notifier: notifier}
}
