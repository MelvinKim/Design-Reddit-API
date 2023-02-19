package usecase

import (
	"log"
	"testing"

	"github.com/MelvinKim/Design-Reddit-API/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserUsecase_CreateUser(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=reddit-api port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to postgres: ", err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := NewUserService(*userRepository)

	type args struct {
		firstName string
		lastName  string
		email     string
		password  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case - Successfully create user",
			args: args{
				firstName: "test ",
				lastName:  "test",
				email:     "testdummy@test.com",
				password:  "test",
			},
			wantErr: false,
		},
		{
			name: "Sad case - Fail to create user - Empty firstName",
			args: args{
				firstName: "",
				lastName:  "test",
				email:     "test@test.com",
				password:  "test",
			},
			wantErr: true,
		},
		{
			name: "Sad case - fail to create user - Empty lastName",
			args: args{
				firstName: "test",
				lastName:  "",
				email:     "test@test.com",
				password:  "test",
			},
			wantErr: true,
		},
		{
			name: "Sad case - fail to create user - Empty email",
			args: args{
				firstName: "test",
				lastName:  "test",
				email:     "",
				password:  "test",
			},
			wantErr: true,
		},
		{
			name: "Sad case - fail to create user - Empty password",
			args: args{
				firstName: "test",
				lastName:  "test",
				email:     "test@test.com",
				password:  "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userService.userRepository.CreateUser(tt.args.firstName, tt.args.lastName, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
