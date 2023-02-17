package entity

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string    `gorm:"size:255; not null;" json:"first_name"`
	LastName  string    `gorm:"size:255; not null;" json:"last_name"`
	Email     string    `gorm:"unique; size:100; not null;" json:"email"`
	Password  string    `gorm:"size:100; not null;" json:"password"`
	IsDeleted bool      `gorm:"column:is_deleted; default:false;" json:"is_deleted"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, err
		default:
			return false, err
		}
	}

	return true, nil
}
