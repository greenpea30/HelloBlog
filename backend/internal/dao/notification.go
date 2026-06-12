package dao

import (
	"helloblog/internal/dao/model"

	"gorm.io/gorm"
)

type NotificationDAO struct {
	db *gorm.DB
}

func NewNotificationDAO(db *gorm.DB) *NotificationDAO {
	return &NotificationDAO{db: db}
}

func (d *NotificationDAO) Create(n *model.Notification) error {
	return d.db.Create(n).Error
}

func (d *NotificationDAO) ListByUser(userID int64, limit int) ([]model.Notification, error) {
	var list []model.Notification
	err := d.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Find(&list).Error
	return list, err
}

func (d *NotificationDAO) UnreadCount(userID int64) (int64, error) {
	var count int64
	err := d.db.Model(&model.Notification{}).
		Where("user_id = ? AND is_read = ?", userID, false).
		Count(&count).Error
	return count, err
}

func (d *NotificationDAO) MarkAllRead(userID int64) error {
	return d.db.Model(&model.Notification{}).
		Where("user_id = ? AND is_read = ?", userID, false).
		Update("is_read", true).Error
}
