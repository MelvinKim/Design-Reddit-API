package usecase

import (
	"errors"

	"github.com/MelvinKim/Design-Reddit-API/entity"
	"github.com/MelvinKim/Design-Reddit-API/repository"
)

type PostService struct {
	postRepository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) *PostService {
	return &PostService{postRepository: postRepository}
}

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

func (s *PostService) GetPost(id int) (*entity.Post, error) {
	if id == 0 {
		return nil, errors.New("please enter the postID")
	}

	return s.postRepository.GetPost(id)
}

func (s *PostService) ListPosts() ([]*entity.Post, error) {
	return s.postRepository.ListPosts()
}
