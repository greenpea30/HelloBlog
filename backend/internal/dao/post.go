package dao

import (
	"sync"
	"time"

	"helloblog/internal/dao/model"

	"gorm.io/gorm"
)

type PostDAO struct {
	db        *gorm.DB
	viewCache map[int64]time.Time
	mu        sync.Mutex
}

func NewPostDAO(db *gorm.DB) *PostDAO {
	return &PostDAO{
		db:        db,
		viewCache: make(map[int64]time.Time),
	}
}

func (d *PostDAO) Create(post *model.Post) (*model.Post, error) {
	err := d.db.Create(post).Error
	return post, err
}

func (d *PostDAO) GetByID(id int64) (*model.Post, error) {
	var post model.Post
	err := d.db.Preload("User").First(&post, "id = ? AND status = ?", id, "normal").Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (d *PostDAO) Update(post *model.Post) error {
	return d.db.Save(post).Error
}

func (d *PostDAO) SoftDelete(id int64, userID int64) error {
	return d.db.Model(&model.Post{}).
		Where("id = ? AND user_id = ?", id, userID).
		Update("status", "deleted").Error
}

type PostListParams struct {
	Page     int
	PageSize int
	UserID   int64
	OrderBy  string
	ZJUOnly  bool
}

func (d *PostDAO) List(params PostListParams) ([]model.Post, int64, error) {
	var posts []model.Post
	var total int64

	query := d.db.Model(&model.Post{}).Where("status = ?", "normal")

	if params.UserID > 0 {
		query = query.Where("user_id = ?", params.UserID)
	}

	if params.ZJUOnly {
		query = query.Where("user_id IN (SELECT id FROM users WHERE zju_id IS NOT NULL)")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	orderBy := "created_at DESC"
	if params.OrderBy == "popular" {
		orderBy = "like_count DESC, comment_count DESC, created_at DESC"
	}

	offset := (params.Page - 1) * params.PageSize
	err := query.Preload("User").
		Order(orderBy).
		Offset(offset).
		Limit(params.PageSize).
		Find(&posts).Error

	return posts, total, err
}

// IncrementView 增加浏览量。返回 true 表示实际执行了 +1，false 表示被缓存拦截
func (d *PostDAO) IncrementView(id int64) (bool, error) {
	d.mu.Lock()
	if expiresAt, ok := d.viewCache[id]; ok && time.Now().Before(expiresAt) {
		d.mu.Unlock()
		return false, nil
	}
	d.viewCache[id] = time.Now().Add(10 * time.Second)
	d.mu.Unlock()

	err := d.db.Model(&model.Post{}).Where("id = ?", id).
		UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error
	return true, err
}

func (d *PostDAO) IncrementLikeCount(id int64) error {
	return d.db.Model(&model.Post{}).Where("id = ?", id).
		UpdateColumn("like_count", gorm.Expr("like_count + 1")).Error
}

func (d *PostDAO) DecrementLikeCount(id int64) error {
	return d.db.Model(&model.Post{}).Where("id = ? AND like_count > 0", id).
		UpdateColumn("like_count", gorm.Expr("like_count - 1")).Error
}

func (d *PostDAO) IncrementCommentCount(id int64) error {
	return d.db.Model(&model.Post{}).Where("id = ?", id).
		UpdateColumn("comment_count", gorm.Expr("comment_count + 1")).Error
}

func (d *PostDAO) ListByUserAndFolder(userID int64, folderID *int64) ([]model.Post, error) {
	var posts []model.Post
	query := d.db.Where("user_id = ? AND status = ?", userID, "normal")
	if folderID == nil {
		query = query.Where("folder_id IS NULL")
	} else {
		query = query.Where("folder_id = ?", *folderID)
	}
	err := query.Order("created_at DESC").Find(&posts).Error
	return posts, err
}

func (d *PostDAO) CountByUserAndFolder(userID int64, folderID *int64) (int64, error) {
	var count int64
	query := d.db.Model(&model.Post{}).Where("user_id = ? AND status = ?", userID, "normal")
	if folderID == nil {
		query = query.Where("folder_id IS NULL")
	} else {
		query = query.Where("folder_id = ?", *folderID)
	}
	err := query.Count(&count).Error
	return count, err
}
