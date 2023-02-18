package repository

import (
	"log"
	"testing"
	"time"

	"github.com/MelvinKim/Design-Reddit-API/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserRepository_CreateUser(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=reddit-api port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to postgres: ", err)
	}

	r := UserRepository{db: db}

	user := &entity.User{
		ID:        49,
		FirstName: "test",
		LastName:  "test",
		Email:     "test@test.com",
		Password:  "test",
		IsDeleted: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := r.db.Create(user).Error; err != nil {
		t.Error("error creating user: ", err)
	}

	retrievedUser, err := r.GetUser(int(user.ID))
	if err != nil {
		t.Error("error retrieving user with ID 49 from the DB: ", err)
	}

	if retrievedUser.ID != user.ID {
		t.Error("Expected ID to be ", user.ID, "but got ", retrievedUser.ID)
	}
	if retrievedUser.FirstName != user.FirstName {
		t.Error("Expected FirstName to be ", user.FirstName, "but got ", retrievedUser.FirstName)
	}
	if retrievedUser.LastName != user.LastName {
		t.Error("Expected LastName to be ", user.LastName, "but got ", retrievedUser.LastName)
	}
	if retrievedUser.Email != user.Email {
		t.Error("Expected Email to be ", user.Email, "but got ", retrievedUser.Email)
	}

	// clean up
	err = r.DeleteUser(int(user.ID))
	if err != nil {
		t.Errorf("Error deleting test user with ID 49 from the DB")
	}
}
