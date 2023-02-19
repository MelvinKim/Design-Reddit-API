package controller

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	usecase "github.com/MelvinKim/Design-Reddit-API/Usecase"
	"github.com/MelvinKim/Design-Reddit-API/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserController_CreateUser(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=reddit-api port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to postgres: ", err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := usecase.NewUserService(*userRepository)
	userController := NewUserController(*userService)

	router := gin.Default()

	router.POST("/users", userController.CreateUser)

	reqBody := []byte(`{"first_name": "test again", "last_name": "test again", "email": "test1@test.com", "password": "test"}`)
	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Error("error creating CreateUser request handler: ", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}
