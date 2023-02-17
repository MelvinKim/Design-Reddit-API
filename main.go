package main

import (
	"log"

	usecase "github.com/MelvinKim/Design-Reddit-API/Usecase"
	"github.com/MelvinKim/Design-Reddit-API/controller"
	"github.com/MelvinKim/Design-Reddit-API/entity"
	"github.com/MelvinKim/Design-Reddit-API/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=reddit-api port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to postgres: ", err)
	}

	// migrate the schema
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatal("failed to migrate schema: ", err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := usecase.NewUserService(*userRepository)
	userController := controller.NewUserController(*userService)

	r := gin.Default()

	r.GET("/users/:id", userController.GetUser)
	r.GET("/users", userController.ListUsers)
	r.POST("/users", userController.CreateUser)

	if err := r.Run(); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
