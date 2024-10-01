package service

import (
	"github.com/lacion/mygolangproject/models"
	"github.com/lacion/mygolangproject/repository"
)

type Service struct {
	UserReposiotry    repository.UserReposiotry
	AddressRepository repository.AddressRepository
}

func NewService(userRepo repository.UserReposiotry, addressRepo repository.AddressRepository) *Service {
	return &Service{
		UserReposiotry:    userRepo,
		AddressRepository: addressRepo,
	}
}

func (s *Service) GetUser(id string) (*models.User, error) {
	user, err := s.UserReposiotry.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
