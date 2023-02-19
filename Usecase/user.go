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
	if first_name == "" {
		return nil, errors.New("please enter your first name")
	}
	if last_name == "" {
		return nil, errors.New("please enter your last name")
	}
	if email == "" {
		return nil, errors.New("please enter your email address")
	}
	if password == "" {
		return nil, errors.New("please enter your password")
	}

	return s.userRepository.CreateUser(first_name, last_name, email, password)
}

func (s *UserService) GetUser(id int) (*entity.User, error) {
	if id == 0 {
		return nil, errors.New("please enter your userID")
	}

	return s.userRepository.GetUser(id)
}

func (s *UserService) ListUsers() ([]*entity.User, error) {
	return s.userRepository.ListUsers()
}
