package entity

import "time"

type Post struct {
	ID            int       `json:"id"`
	Creator       int       `json:"creator_id"` // TODO: check if to use an int or a whole struct to reference the creator
	Subreddit     int       `json:"subrredit_id"`
	Title         string    `json:"string"`
	Content       string    `json:"content"`
	VotesCount    int       `json:"votes_count"`
	CommentsCount int       `json:"comments_count"`
	IsDeleted     bool      `json:"is_deleted"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
