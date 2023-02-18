package repository

import (
	"time"

	"github.com/MelvinKim/Design-Reddit-API/entity"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) CreateComment(userID, postID int, content string) (*entity.Comment, error) {
	comment := &entity.Comment{
		Creator:    userID,
		Post:       postID,
		Content:    content,
		VotesCount: 0,
		IsDeleted:  false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err := r.db.Create(comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *CommentRepository) ListComments(userID, postID int) ([]*entity.Comment, error) {
	comments := []*entity.Comment{}
	if err := r.db.Where("creator = ? AND post = ?", userID, postID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
