package usecase

import (
	"log"
	"testing"

	"github.com/MelvinKim/Design-Reddit-API/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestPostUsecase_CreatePost(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=reddit-api port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to postgres: ", err)
	}

	postRepository := repository.NewPostRepository(db)
	postService := NewPostService(*postRepository)

	type args struct {
		title     string
		content   string
		creator   int
		subreddit int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case - Successfully create post",
			args: args{
				title:     "test",
				content:   "test",
				creator:   45,
				subreddit: 45,
			},
			wantErr: false,
		},
		{
			name: "Sad case - fail to create post - Empty title",
			args: args{
				title:     "",
				content:   "test",
				creator:   45,
				subreddit: 45,
			},
			wantErr: true,
		},
		{
			name: "Sad case - fail to create post - Empty post body",
			args: args{
				title:     "test",
				content:   "",
				creator:   45,
				subreddit: 45,
			},
			wantErr: true,
		},
		{
			name: "Sad case - fail to create post - Empty creatorID",
			args: args{
				title:     "test",
				content:   "test",
				creator:   0,
				subreddit: 45,
			},
			wantErr: true,
		},
		{
			name: "Sad case - fail to create post - Empty subredditID",
			args: args{
				title:     "test",
				content:   "test",
				creator:   45,
				subreddit: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			post, err := postService.CreatePost(tt.args.creator, tt.args.subreddit, tt.args.title, tt.args.content)
			if post == nil && !tt.wantErr {
				t.Errorf("PostService.CreatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("PostService.CreatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
