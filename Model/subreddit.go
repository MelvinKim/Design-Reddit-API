package model

import "time"

type Subreddit struct {
	ID        int       `json:"id"`
	Creator   User      `json:"creator_id"` // TODO: check if to use an int or a whole struct to reference the creator
	Name      int       `json:"name"`
	Redditors []User    `json:"redditors"`
	Posts     []Post    `json:"posts"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
