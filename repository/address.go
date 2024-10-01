package repository

import (
	"github.com/lacion/mygolangproject/models"
	"gorm.io/gorm"
)

type AddressRepository interface {
	CreateAddress(address *models.Address) error
}

type AddressRepositoryImpl struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &AddressRepositoryImpl{db: db}
}

func (a *AddressRepositoryImpl) CreateAddress(address *models.Address) error {
	return a.db.Create(address).Error
}
