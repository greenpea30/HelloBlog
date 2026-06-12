package dao

import (
	"helloblog/internal/dao/model"

	"gorm.io/gorm"
)

type LinkDAO struct {
	db *gorm.DB
}

func NewLinkDAO(db *gorm.DB) *LinkDAO {
	return &LinkDAO{db: db}
}

func (d *LinkDAO) Create(link *model.Link) error {
	return d.db.Create(link).Error
}

func (d *LinkDAO) List() ([]model.Link, error) {
	var links []model.Link
	err := d.db.Order("created_at DESC").Find(&links).Error
	return links, err
}

func (d *LinkDAO) Delete(id int64) error {
	return d.db.Delete(&model.Link{}, id).Error
}
