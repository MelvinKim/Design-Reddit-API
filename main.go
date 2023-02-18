package main

import (
	"log"
	"net/http"

	usecase "github.com/MelvinKim/Design-Reddit-API/Usecase"
	"github.com/MelvinKim/Design-Reddit-API/controller"
	"github.com/MelvinKim/Design-Reddit-API/entity"
	"github.com/MelvinKim/Design-Reddit-API/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func HomepageHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to the Reddit API build with Golang"})
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=reddit-api port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to postgres: ", err)
	}

	// migrate the schema
	err = db.AutoMigrate(&entity.User{}, &entity.Post{}, &entity.Comment{})
	if err != nil {
		log.Fatal("failed to migrate schema: ", err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := usecase.NewUserService(*userRepository)
	userController := controller.NewUserController(*userService)

	postRepository := repository.NewPostRepository(db)
	postService := usecase.NewPostService(*postRepository)
	postController := controller.NewPostController(*postService)

	commentRepository := repository.NewCommentRepository(db)
	commentService := usecase.NewCommentService(*commentRepository)
	commentController := controller.NewCommentController(*commentService)

	r := gin.Default()

	r.GET("/", HomepageHandler)

	r.GET("/users/:id", userController.GetUser)
	r.GET("/users", userController.ListUsers)
	r.POST("/users", userController.CreateUser)

	r.GET("/posts/:id", postController.GetPost)
	r.GET("/posts", postController.ListPosts)
	r.POST("/posts", postController.CreatePost)

	r.POST("/comments", commentController.CreateComment)
	r.GET("/comments/:userID/:postID", commentController.ListComments)

	if err := r.Run(); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
