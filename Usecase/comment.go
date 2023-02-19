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

func (s *CommentService) ListComments(userID, postID int) ([]*entity.Comment, error) {
	if userID == 0 {
		return nil, errors.New("please enter a userID")
	}
	if postID == 0 {
		return nil, errors.New("please enter a postID")
	}

	return s.commentRepository.ListComments(userID, postID)
}
