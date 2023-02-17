package repository

import (
	"time"

	"github.com/MelvinKim/Design-Reddit-API/entity"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(creator, subreddit int, title, content string) (*entity.Post, error) {
	post := &entity.Post{
		Creator:       creator,
		Subreddit:     subreddit,
		Title:         title,
		Content:       content,
		VotesCount:    0,
		CommentsCount: 0,
		IsDeleted:     false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := r.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) GetPost(id int) (*entity.Post, error) {
	post := &entity.Post{}
	if err := r.db.First(post, id).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) ListPosts() ([]*entity.Post, error) {
	posts := []*entity.Post{}
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
