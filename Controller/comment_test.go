package controller

import (
	"bytes"
	"fmt"
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

func TestComment_CreateComment(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=reddit-api port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to postgres: ", err)
	}

	commentRepository := repository.NewCommentRepository(db)
	commentService := usecase.NewCommentService(*commentRepository)
	commentController := NewCommentController(*commentService)

	router := gin.Default()

	router.POST("/comments", commentController.CreateComment)

	reqBody := []byte(`{"creator_id": 1, "post_id": 1, "content": "test comment"}`)
	req, err := http.NewRequest(http.MethodPost, "/comments", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Error("error creating CreateComment request handler: ", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	fmt.Println("status code: ", rr.Code)
	assert.Equal(t, http.StatusCreated, rr.Code)
}
