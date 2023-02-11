package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	model "github.com/MelvinKim/Design-Reddit-API/Controller"
	controller "github.com/MelvinKim/Design-Reddit-API/Controller"
	postgres "github.com/MelvinKim/Design-Reddit-API/Repository/postgres"
	router "github.com/MelvinKim/Design-Reddit-API/Router"
	usecase "github.com/MelvinKim/Design-Reddit-API/Usecase"
	"github.com/MelvinKim/Design-Reddit-API/repository"
	"github.com/jinzhu/gorm"
)

var (
	DB             *gorm.DB
	userRepository repository.UserRepository = postgres.NewPostgresRepository(DB)
	userUsecase    usecase.UserUsecase       = usecase.NewUserUseCase(userRepository)
	userController controller.UserController = controller.NewUserController(userUsecase)
	httpRouter     router.Router             = router.NewChiRouter()
)

type Server struct {
	DB *gorm.DB
}

// create a database connection
func (server *Server) connectToDB(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
	server.DB, err = gorm.Open(dbDriver, DBURL)
	if err != nil {
		log.Fatalf("Error while trying to connect to % DB: %v", err)
	} else {
		log.Default().Printf("Connected to %s Postgres DB successfully!!", dbName)
	}

	// database migration
	server.DB.Debug().AutoMigrate(model.User{}, )) 

}

func main() {
	const port string = ":8080"

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
