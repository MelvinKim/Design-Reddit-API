package usecase

import (
	"log"
	"testing"

	"github.com/MelvinKim/Design-Reddit-API/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCommentUsecase_CreateComment(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=reddit-api port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to postgres: ", err)
	}

	commentRepository := repository.NewCommentRepository(db)
	commentService := NewCommentService(*commentRepository)

	type args struct {
		userID  int
		postID  int
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case - Successfully create comment",
			args: args{
				userID:  45,
				postID:  45,
				content: "test comment",
			},
			wantErr: false,
		},
		{
			name: "Sad case - Fail to create comment - Empty userID",
			args: args{
				userID:  0,
				postID:  45,
				content: "test comment",
			},
			wantErr: true,
		},
		{
			name: "Sad case - Fail to create comment - Empty postID",
			args: args{
				userID:  45,
				postID:  0,
				content: "test comment",
			},
			wantErr: true,
		},
		{
			name: "Sad case - Fail to create comment - Empty comment body",
			args: args{
				userID:  45,
				postID:  45,
				content: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment, err := commentService.CreateComment(tt.args.userID, tt.args.postID, tt.args.content)
			if comment == nil && !tt.wantErr {
				t.Errorf("CommentService.CreateComment() error = %v, wantErr %v", err, tt.wantErr)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentService.CreateComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
