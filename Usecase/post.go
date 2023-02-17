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
	if title == "" || content == "" || creator == 0 || subreddit == 0 {
		return nil, errors.New("all fields are required")
	}
	return s.postRepository.CreatePost(creator, subreddit, title, content)
}

func (s *PostService) GetPost(id int) (*entity.Post, error) {
	return s.postRepository.GetPost(id)
}

func (s *PostService) ListPosts() ([]*entity.Post, error) {
	return s.postRepository.ListPosts()
}
