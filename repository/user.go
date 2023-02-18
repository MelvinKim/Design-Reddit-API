package repository

import (
	"time"

	"github.com/MelvinKim/Design-Reddit-API/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(first_name, last_name, email, password string) (*entity.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &entity.User{
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
		Password:  string(hashPassword),
		IsDeleted: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUser(id int) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) ListUsers() ([]*entity.User, error) {
	users := []*entity.User{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
