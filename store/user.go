package store

import (
	"company-keeper-go/models"
	"errors"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) GetByID(id uint) (*models.User, error) {
	var m models.User
	if err := us.db.First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) GetByEmail(e string) (*models.User, error) {
	var m models.User
	if err := us.db.Where(&models.User{Email: e}).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) Create(u *models.User) (err error) {
	return us.db.Create(u).Error
}

func (us *UserStore) Delete(u *models.User) error {
	return us.db.Model(u).Delete(u).Error
}
