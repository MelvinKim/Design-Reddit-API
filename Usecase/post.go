package usecase

import (
	"errors"

	"github.com/MelvinKim/Design-Reddit-API/entity"
	"github.com/MelvinKim/Design-Reddit-API/repository"
)

// PostService is responsible for implementing all post related business logic
type PostService struct {
	postRepository repository.PostRepository
}

// NewPostService creates a new PostService Instance
func NewPostService(postRepository repository.PostRepository) *PostService {
	return &PostService{postRepository: postRepository}
}

// CreatePost handles the post creation logic
func (s *PostService) CreatePost(creator, subreddit int, title, content string) (*entity.Post, error) {
	if title == "" {
		return nil, errors.New("please enter a post title")
	}
	if content == "" {
		return nil, errors.New("please populate your post :)")
	}
	if creator == 0 {
		return nil, errors.New("please provide the creatorID")
	}
	if subreddit == 0 {
		return nil, errors.New("please provide the subredditID")
	}

	return s.postRepository.CreatePost(creator, subreddit, title, content)
}

// GetPost handles the logic for fetching posts using postIDs
func (s *PostService) GetPost(id int) (*entity.Post, error) {
	if id == 0 {
		return nil, errors.New("please enter the postID")
	}

	return s.postRepository.GetPost(id)
}

// ListPosts handles the logic for fetching all application posts
func (s *PostService) ListPosts() ([]*entity.Post, error) {
	return s.postRepository.ListPosts()
}
