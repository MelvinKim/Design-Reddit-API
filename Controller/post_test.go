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

func TestPostController_CreatePost(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=reddit-api port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to postgres: ", err)
	}

	postRepository := repository.NewPostRepository(db)
	postService := usecase.NewPostService(*postRepository)
	postController := NewPostController(*postService)

	router := gin.Default()

	router.POST("/posts", postController.CreatePost)

	reqBody := []byte(`{"creator_id": 4, "subrredit_id": 4, "title": "test again", "content": "test again"}`)
	req, err := http.NewRequest(http.MethodPost, "/posts", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Error("error creating CreatePost request handler: ", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}
