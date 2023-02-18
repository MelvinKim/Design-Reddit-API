package repository

import (
	"log"
	"testing"
	"time"

	"github.com/MelvinKim/Design-Reddit-API/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCommentRepository_CreateComment(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=reddit-api port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to postgres: ", err)
	}

	r := UserRepository{db: db}

	comment := &entity.Comment{
		ID:         49,
		Creator:    49,
		Post:       49,
		VotesCount: 0,
		Content:    "test",
		IsDeleted:  false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err := r.db.Create(comment).Error; err != nil {
		t.Error("error creating comment: ", err)
	}

	retrievedComment, err := r.GetComment(int(comment.ID))
	if err != nil {
		t.Error("error retrieving comment with ID 49 from the DB: ", err)
	}

	if retrievedComment.ID != comment.ID {
		t.Error("Expected Comment ID to be ", comment.ID, "but got ", retrievedComment.ID)
	}
	if retrievedComment.Creator != comment.Creator {
		t.Error("Expected Comment Creator to be ", comment.ID, "but got ", retrievedComment.Creator)
	}
	if retrievedComment.Post != comment.Post {
		t.Error("Expected Comment Post to be ", comment.ID, "but got ", retrievedComment.Post)
	}
	if retrievedComment.Content != comment.Content {
		t.Error("Expected Comment content to be ", comment.ID, "but got ", retrievedComment.Content)
	}

	// clean up
	err = r.DeleteComment(comment.ID)
	if err != nil {
		t.Errorf("Error deleting test comment with ID 49 from the DB")
	}
}
