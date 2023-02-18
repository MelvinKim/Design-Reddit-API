package usecase

import (
	"errors"

	"github.com/MelvinKim/Design-Reddit-API/entity"
	"github.com/MelvinKim/Design-Reddit-API/repository"
)

type CommentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(commentRepository repository.CommentRepository) *CommentService {
	return &CommentService{commentRepository: commentRepository}
}

func (s *CommentService) CreateComment(userID, postID int, content string) (*entity.Comment, error) {
	if userID == 0 || postID == 0 || content == "" {
		return nil, errors.New("all fields are required")
	}

	return s.commentRepository.CreateComment(userID, postID, content)
}

func (s *CommentService) ListComments(userID, postID int) ([]*entity.Comment, error) {
	if userID == 0 || postID == 0 {
		return nil, errors.New("all fields are required")
	}

	return s.commentRepository.ListComments(userID, postID)
}
