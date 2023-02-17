package entity

import "time"

type Subreddit struct {
	ID        int       `json:"id"`
	Creator   int       `json:"creator_id"` // TODO: check if to use an int or a whole struct to reference the creator
	Name      int       `json:"name"`
	Redditors []int     `json:"redditors"`
	Posts     []int     `json:"posts"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
