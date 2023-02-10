package main

import (
	"fmt"
	"net/http"

	router "github.com/MelvinKim/Design-Reddit-API/http"
)

var (
	httpRouter router.Router = router.NewChiRouter()
)

func main() {
	const port string = ":8080"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "server is app and running ....")
	})

	// TODO: register the route groups here

	httpRouter.SERVE(port)

}
