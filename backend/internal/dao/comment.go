package dao

import (
	"helloblog/internal/dao/model"

	"gorm.io/gorm"
)

type CommentDAO struct {
	db *gorm.DB
}

func NewCommentDAO(db *gorm.DB) *CommentDAO {
	return &CommentDAO{db: db}
}

func (d *CommentDAO) Create(comment *model.Comment) (*model.Comment, error) {
	err := d.db.Create(comment).Error
	return comment, err
}

func (d *CommentDAO) GetByID(id int64) (*model.Comment, error) {
	var comment model.Comment
	err := d.db.Preload("User").Preload("Children.User").
		First(&comment, "id = ? AND status = ?", id, "normal").Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (d *CommentDAO) ListByPost(postID int64, parentID *int64) ([]model.Comment, error) {
	var comments []model.Comment
	query := d.db.Preload("User").Preload("Children.User").
		Where("post_id = ? AND status = ?", postID, "normal")

	if parentID == nil {
		query = query.Where("parent_id IS NULL")
	} else {
		query = query.Where("parent_id = ?", *parentID)
	}

	err := query.Order("created_at ASC").Find(&comments).Error
	return comments, err
}

func (d *CommentDAO) SoftDelete(id int64, userID int64) error {
	return d.db.Model(&model.Comment{}).
		Where("id = ? AND user_id = ?", id, userID).
		Update("status", "deleted").Error
}

func (d *CommentDAO) IncrementLikeCount(id int64) error {
	return d.db.Model(&model.Comment{}).Where("id = ?", id).
		UpdateColumn("like_count", gorm.Expr("like_count + 1")).Error
}

func (d *CommentDAO) DecrementLikeCount(id int64) error {
	return d.db.Model(&model.Comment{}).Where("id = ? AND like_count > 0", id).
		UpdateColumn("like_count", gorm.Expr("like_count - 1")).Error
}
