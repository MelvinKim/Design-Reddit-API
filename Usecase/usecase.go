package usecase

import model "github.com/MelvinKim/Design-Reddit-API/Model"

type UserUsecase interface {
	Create(user *model.User) (*model.User, error)
	Get(id int) (*model.User, error)
	Update(id int, user *model.User) (*model.User, error)
	FindAll() ([]*model.User, error)
	Delete(id int) (int64, error)
}

type SubredditUsecase interface {
	Create(subrredit *model.Subreddit) (*model.Subreddit, error)
	Get(id int) (*model.Subreddit, error)
	Update(id int, subreddit *model.Subreddit) (*model.Subreddit, error)
	Delete(id int) (int64, error)
	FindAll() ([]*model.Subreddit, error)
}

type PostUsecase interface {
	Create(post *model.Post) (*model.Post, error)
	Get(id int) (*model.Post, error)
	Update(id int, post *model.Post) (*model.Post, error)
	FindAll() ([]*model.Post, error)
	Delete(id int) (int64, error)
}

type CommentUsecase interface {
	Create(comment *model.Comment) (*model.Comment, error)
	Get(id int) (*model.Comment, error)
	Update(id int, comment *model.Comment) (*model.Comment, error)
	FindAll() ([]*model.Comment, error)
	Delete(id int) (int64, error)
}
