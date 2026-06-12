package dao

import (
	"errors"

	"helloblog/internal/dao/model"
	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

var ErrNotFound = gorm.ErrRecordNotFound

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func (d *UserDAO) Create(user *model.User) error {
	return d.db.Create(user).Error
}

func (d *UserDAO) GetByID(id int64) (*model.User, error) {
	var user model.User
	if err := d.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *UserDAO) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := d.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *UserDAO) GetByZjuID(zjuID string) (*model.User, error) {
	var user model.User
	if err := d.db.Where("zju_id = ?", zjuID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *UserDAO) UsernameExists(username string) (bool, error) {
	var count int64
	err := d.db.Model(&model.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

func (d *UserDAO) EmailExists(email string) (bool, error) {
	var count int64
	err := d.db.Model(&model.User{}).Where("email = ? AND email IS NOT NULL", email).Count(&count).Error
	return count > 0, err
}

func (d *UserDAO) ZjuIDExists(zjuID string) (bool, error) {
	var count int64
	err := d.db.Model(&model.User{}).Where("zju_id = ?", zjuID).Count(&count).Error
	return count > 0, err
}

func (d *UserDAO) Update(user *model.User) error {
	return d.db.Save(user).Error
}
