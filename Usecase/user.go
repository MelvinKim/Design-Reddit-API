package usecase

import (
	model "github.com/MelvinKim/Design-Reddit-API/Model"
	"github.com/MelvinKim/Design-Reddit-API/repository"
)

type userUsecase struct{}

var (
	userRepository repository.UserRepository
)

func NewUserUseCase(repository repository.UserRepository) UserUsecase {
	userRepository = repository
	return &userUsecase{}
}

func (u *userUsecase) Create(user *model.User) (*model.User, error) {
	return userRepository.Create(user)
}

func (u *userUsecase) Get(id int) (*model.User, error) {
	return userRepository.Get(id)
}

func (u *userUsecase) Update(id int, user *model.User) (*model.User, error) {
	return userRepository.Update(id, user)
}

func (u *userUsecase) FindAll() ([]*model.User, error) {
	return userRepository.FindAll()
}

func (u *userUsecase) Delete(id int) (int64, error) {
	return userRepository.Delete(id)
}
