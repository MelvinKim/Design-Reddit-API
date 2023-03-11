package repository

import (
	"time"

	"github.com/MelvinKim/Design-Reddit-API/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepository hanldes the Database logic for the User entity
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser creates a new user in the DB
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

// GetUser fetches a user from the DB based on the userID
func (r *UserRepository) GetUser(id int) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// ListUsers fetches all users from the DB
func (r *UserRepository) ListUsers() ([]*entity.User, error) {
	users := []*entity.User{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// DeleteUser deletes a User from the DB, based on the UserID
func (r *UserRepository) DeleteUser(id int) error {
	user := &entity.User{}
	if err := r.db.Delete(user, id).Error; err != nil {
		return err
	}

	return nil
}
