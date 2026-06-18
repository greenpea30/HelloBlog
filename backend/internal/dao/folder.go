package dao

import (
	"helloblog/internal/dao/model"

	"gorm.io/gorm"
)

type FolderDAO struct {
	db *gorm.DB
}

func NewFolderDAO(db *gorm.DB) *FolderDAO {
	return &FolderDAO{db: db}
}

func (d *FolderDAO) Create(folder *model.Folder) error {
	return d.db.Create(folder).Error
}

func (d *FolderDAO) GetByID(id int64) (*model.Folder, error) {
	var folder model.Folder
	err := d.db.First(&folder, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &folder, nil
}

func (d *FolderDAO) ListByUser(userID int64) ([]model.Folder, error) {
	var folders []model.Folder
	err := d.db.Where("user_id = ?", userID).
		Order("created_at ASC").
		Find(&folders).Error
	return folders, err
}

func (d *FolderDAO) Delete(id int64, userID int64) error {
	return d.db.Where("id = ? AND user_id = ?", id, userID).
		Delete(&model.Folder{}).Error
}

func (d *FolderDAO) Exists(id int64, userID int64) (bool, error) {
	var count int64
	err := d.db.Model(&model.Folder{}).
		Where("id = ? AND user_id = ?", id, userID).
		Count(&count).Error
	return count > 0, err
}

// ClearFolderID 删除文件夹时将其下文章的 folder_id 置空
func ClearFolderID(db *gorm.DB, folderID int64) error {
	return db.Model(&model.Post{}).
		Where("folder_id = ?", folderID).
		Update("folder_id", nil).Error
}
