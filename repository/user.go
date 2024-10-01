package repository

import (
	"github.com/lacion/mygolangproject/models"
	"gorm.io/gorm"
)

type UserReposiotry interface {
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) error
	GetCount() (int64, error)
}

type UserReposiotryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserReposiotry {
	return &UserReposiotryImpl{db: db}
}

func (u *UserReposiotryImpl) GetUserByID(id string) (*models.User, error) {
	user := &models.User{}
	err := u.db.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserReposiotryImpl) CreateUser(user *models.User) error {
	return u.db.Create(user).Error
}

func (u *UserReposiotryImpl) GetCount() (int64, error) {
	var count int64
	err := u.db.Model(&models.User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
