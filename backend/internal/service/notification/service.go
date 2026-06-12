package notification

import (
	"helloblog/internal/dao/model"
	"helloblog/internal/dto"
)

type UseCase interface {
	Create(userID int64, nType, title, content string, fromUserID, postID *int64) error
	List(userID int64) ([]dto.NotificationResponse, error)
	UnreadCount(userID int64) (int64, error)
	MarkAllRead(userID int64) error
}

type notificationRepository interface {
	Create(n *model.Notification) error
	ListByUser(userID int64, limit int) ([]model.Notification, error)
	UnreadCount(userID int64) (int64, error)
	MarkAllRead(userID int64) error
}

type Service struct {
	repo notificationRepository
}

func NewService(repo notificationRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(userID int64, nType, title, content string, fromUserID, postID *int64) error {
	n := &model.Notification{
		UserID:     userID,
		Type:       nType,
		Title:      title,
		Content:    content,
		FromUserID: fromUserID,
		PostID:     postID,
	}
	return s.repo.Create(n)
}

func (s *Service) List(userID int64) ([]dto.NotificationResponse, error) {
	list, err := s.repo.ListByUser(userID, 50)
	if err != nil {
		return nil, err
	}
	resp := make([]dto.NotificationResponse, len(list))
	for i, n := range list {
		resp[i] = dto.NotificationResponse{
			ID:         n.ID,
			Type:       n.Type,
			Title:      n.Title,
			Content:    n.Content,
			FromUserID: n.FromUserID,
			PostID:     n.PostID,
			IsRead:     n.IsRead,
			CreatedAt:  n.CreatedAt,
		}
	}
	return resp, nil
}

func (s *Service) UnreadCount(userID int64) (int64, error) {
	return s.repo.UnreadCount(userID)
}

func (s *Service) MarkAllRead(userID int64) error {
	return s.repo.MarkAllRead(userID)
}
