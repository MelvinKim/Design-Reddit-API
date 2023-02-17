package usecase

import (
	"errors"

	"github.com/MelvinKim/Design-Reddit-API/entity"
	"github.com/MelvinKim/Design-Reddit-API/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(first_name, last_name, email, password string) (*entity.User, error) {
	if first_name == "" || last_name == "" || email == "" || password == "" {
		return nil, errors.New("please provide all the fields")
	}
	return s.userRepository.CreateUser(first_name, last_name, email, password)
}

func (s *UserService) GetUser(id int) (*entity.User, error) {
	return s.userRepository.GetUser(id)
}

func (s *UserService) ListUsers() ([]*entity.User, error) {
	return s.userRepository.ListUsers()
}
