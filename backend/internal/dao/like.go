package dao

import (
	"helloblog/internal/dao/model"

	"gorm.io/gorm"
)

type LikeDAO struct {
	db *gorm.DB
}

func NewLikeDAO(db *gorm.DB) *LikeDAO {
	return &LikeDAO{db: db}
}

// Toggle 切换点赞状态，返回是否已点赞
func (d *LikeDAO) Toggle(userID int64, targetType string, targetID int64) (liked bool, err error) {
	var existing model.Like
	err = d.db.Where("user_id = ? AND target_type = ? AND target_id = ?",
		userID, targetType, targetID).First(&existing).Error

	if err == nil {
		// 已点赞，取消
		if delErr := d.db.Delete(&existing).Error; delErr != nil {
			return false, delErr
		}
		return false, nil
	}

	if !IsNotFound(err) {
		return false, err
	}

	// 未点赞，创建
	like := &model.Like{
		UserID:     userID,
		TargetType: targetType,
		TargetID:   targetID,
	}
	if createErr := d.db.Create(like).Error; createErr != nil {
		return false, createErr
	}
	return true, nil
}

func (d *LikeDAO) IsLiked(userID int64, targetType string, targetID int64) (bool, error) {
	var count int64
	err := d.db.Model(&model.Like{}).
		Where("user_id = ? AND target_type = ? AND target_id = ?", userID, targetType, targetID).
		Count(&count).Error
	return count > 0, err
}

func (d *LikeDAO) GetUserLikedPostIDs(userID int64) ([]int64, error) {
	var ids []int64
	err := d.db.Model(&model.Like{}).
		Where("user_id = ? AND target_type = ?", userID, "post").
		Pluck("target_id", &ids).Error
	return ids, err
}
