package user_service

import (
	"cleara/internal/core/domain"
	"cleara/internal/core/ports"
)

type UserService struct {
	usersRepository ports.UserRepository
}

func New(repository ports.UserRepository) *UserService {
	return &UserService{
		usersRepository: repository,
	}
}

func (srv *UserService) GetProfile(id string) (*domain.User, error) {
	userProfile, err := srv.usersRepository.GetProfile(id)
	if err != nil {
		return nil, err
	}
	return userProfile, nil
}
