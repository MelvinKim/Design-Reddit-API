package model

import "time"

type Comment struct {
	ID         int       `json:"int"`
	Creator    User      `json:"creator_id"`
	Post       Post      `json:"post_id"`
	VotesCount int       `json:"votes_count"`
	Content    string    `json:"content"`
	IsDeleted  bool      `json:"is_deleted"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
