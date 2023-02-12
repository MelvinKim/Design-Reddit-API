package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	controller "github.com/MelvinKim/Design-Reddit-API/Controller"
	postgres "github.com/MelvinKim/Design-Reddit-API/Repository/postgres"
	router "github.com/MelvinKim/Design-Reddit-API/Router"
	usecase "github.com/MelvinKim/Design-Reddit-API/Usecase"
	"github.com/MelvinKim/Design-Reddit-API/repository"
	"github.com/jackc/pgx/v4"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	DB             *gorm.DB
	DSN            string
	userRepository repository.UserRepository = postgres.NewPostgresRepository(DB)
	userUsecase    usecase.UserUsecase       = usecase.NewUserUseCase(userRepository)
	userController controller.UserController = controller.NewUserController(userUsecase)
	httpRouter     router.Router             = router.NewChiRouter()
)

// create a database connection
func connectToDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	DSN = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	dbConn, err := pgx.Connect(context.Background(), DSN)
	if err != nil {
		log.Fatalf("unable to connect to postgres :) err: %v\n", err)
	}

	fmt.Println("Connected to postgres successfully!! :)")
	defer dbConn.Close(context.Background())
}

func main() {
	const port string = ":8080"

	// connect to Postgres
	connectToDB()

	// Homepage route
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		message := map[string]string{"status": "server up and running :)"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message)
	})

	// Users routes
	httpRouter.GET("/users/{id}", userController.Get)
	httpRouter.GET("/users", userController.FindAll)
	httpRouter.POST("/users", userController.Create)
	httpRouter.PUT("/users/{id}", userController.Update)
	httpRouter.DELETE("/users/{id}", userController.Delete)

	httpRouter.SERVE(port)
}
