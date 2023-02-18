package repository

import (
	"log"
	"testing"
	"time"

	"github.com/MelvinKim/Design-Reddit-API/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestPostRepository_CreatePost(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=reddit-api port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to postgres: ", err)
	}

	r := PostRepository{db: db}

	post := &entity.Post{
		ID:            49,
		Creator:       49,
		Subreddit:     49,
		Title:         "test",
		Content:       "test",
		VotesCount:    0,
		CommentsCount: 0,
		IsDeleted:     false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := r.db.Create(post).Error; err != nil {
		t.Error("error creating post: ", err)
	}

	retrievedPost, err := r.GetPost(post.ID)
	if err != nil {
		t.Error("error retrieving post with ID 49 from DB: ", err)
	}

	if retrievedPost.ID != post.ID {
		t.Error("Expected Post ID to be ", post.ID, "but got ", retrievedPost.ID)
	}
	if retrievedPost.Content != post.Content {
		t.Error("Expected Post content to be ", post.Content, "but got ", retrievedPost.Content)
	}
	if retrievedPost.Title != post.Title {
		t.Error("Expected Post Title to be ", post.Title, "but got ", retrievedPost.Title)
	}
	if retrievedPost.Creator != post.Creator {
		t.Error("Expected Post Creator to be ", post.Creator, "but got ", retrievedPost.Creator)
	}
	if retrievedPost.Subreddit != post.Subreddit {
		t.Error("Expected Post Subreddit to be ", post.Subreddit, "but got ", retrievedPost.Subreddit)
	}

	// clean up
	err = r.DeletePost(post.ID)
	if err != nil {
		t.Error("error deleting test post with ID 49 from the DB")
	}
}
