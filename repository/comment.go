package repository

import (
	"time"

	"github.com/MelvinKim/Design-Reddit-API/entity"
	"gorm.io/gorm"
)

// CommentRepository hanldes the Database logic for the Comment entity
type CommentRepository struct {
	db *gorm.DB
}

// NewCommentRepository creates a new CommentRepository instance
func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

// CreateComment creates a new comment in the DB
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

// ListComments fetches all comments made by a user on a post from the DB
func (r *CommentRepository) ListComments(userID, postID int) ([]*entity.Comment, error) {
	comments := []*entity.Comment{}
	if err := r.db.Where("creator = ? AND post = ?", userID, postID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

// GetComment fetches a comment based on the commentID
func (r *UserRepository) GetComment(id int) (*entity.Comment, error) {
	comment := &entity.Comment{}
	if err := r.db.First(comment, id).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

// DeleteComment deletes a comment from the DB, using the commentID
func (r *UserRepository) DeleteComment(id int) error {
	comment := &entity.Comment{}
	if err := r.db.Delete(comment, id).Error; err != nil {
		return err
	}

	return nil
}
