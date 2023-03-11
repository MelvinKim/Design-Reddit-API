package main

import (
	"log"
	"net/http"

	controller "github.com/MelvinKim/Design-Reddit-API/Controller"
	usecase "github.com/MelvinKim/Design-Reddit-API/Usecase"
	"github.com/MelvinKim/Design-Reddit-API/entity"
	"github.com/MelvinKim/Design-Reddit-API/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	UserController    *controller.UserController
	PostController    *controller.PostControler
	CommentController *controller.CommentController
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
	UserController := controller.NewUserController(*userService)

	postRepository := repository.NewPostRepository(db)
	postService := usecase.NewPostService(*postRepository)
	PostController := controller.NewPostController(*postService)

	commentRepository := repository.NewCommentRepository(db)
	commentService := usecase.NewCommentService(*commentRepository)
	CommentController := controller.NewCommentController(*commentService)

	r := gin.Default()

	r.GET("/", HomepageHandler)

	r.GET("/users/:id", UserController.GetUser)
	r.GET("/users", UserController.ListUsers)
	r.POST("/users", UserController.CreateUser)

	r.GET("/posts/:id", PostController.GetPost)
	r.GET("/posts", PostController.ListPosts)
	r.POST("/posts", PostController.CreatePost)

	r.POST("/comments", CommentController.CreateComment)
	r.GET("/comments/:userID/:postID", CommentController.ListComments)

	if err := r.Run(); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
