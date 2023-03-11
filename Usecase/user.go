package usecase

import (
	"errors"

	"github.com/MelvinKim/Design-Reddit-API/entity"
	"github.com/MelvinKim/Design-Reddit-API/repository"
)

// UserService is responsible for implementing all user related logic
type UserService struct {
	userRepository repository.UserRepository
}

// NewUserService creates a new UserService Instance
func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

// CreateUser handles the user creation logic
func (s *UserService) CreateUser(firstName, lastName, email, password string) (*entity.User, error) {
	if firstName == "" {
		return nil, errors.New("please enter your first name")
	}
	if lastName == "" {
		return nil, errors.New("please enter your last name")
	}
	if email == "" {
		return nil, errors.New("please enter your email address")
	}
	if password == "" {
		return nil, errors.New("please enter your password")
	}

	return s.userRepository.CreateUser(firstName, lastName, email, password)
}

// GetUser handles all the logic for fetching user using UserIDs
func (s *UserService) GetUser(id int) (*entity.User, error) {
	if id == 0 {
		return nil, errors.New("please enter your userID")
	}

	return s.userRepository.GetUser(id)
}

// ListUsers handles the logic for fetching application users
func (s *UserService) ListUsers() ([]*entity.User, error) {
	return s.userRepository.ListUsers()
}
