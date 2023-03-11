package usecase

import (
	"errors"

	"github.com/MelvinKim/Design-Reddit-API/entity"
	"github.com/MelvinKim/Design-Reddit-API/repository"
)

// CommentService is responsible for implementing all comment related logic
type CommentService struct {
	commentRepository repository.CommentRepository
}

// NewCommentService creates a new CommentService Instance
func NewCommentService(commentRepository repository.CommentRepository) *CommentService {
	return &CommentService{commentRepository: commentRepository}
}

// CreateComment handles the comment creation logic
func (s *CommentService) CreateComment(userID, postID int, content string) (*entity.Comment, error) {
	if userID == 0 {
		return nil, errors.New("please enter a userID")
	}
	if postID == 0 {
		return nil, errors.New("please enter a postID")
	}
	if content == "" {
		return nil, errors.New("please enter the comment body")
	}

	return s.commentRepository.CreateComment(userID, postID, content)
}

// ListComments handles the logic for fetching all comments made by a user on a specific post
func (s *CommentService) ListComments(userID, postID int) ([]*entity.Comment, error) {
	if userID == 0 {
		return nil, errors.New("please enter a userID")
	}
	if postID == 0 {
		return nil, errors.New("please enter a postID")
	}

	return s.commentRepository.ListComments(userID, postID)
}
