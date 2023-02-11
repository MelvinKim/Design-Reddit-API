package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type chiRouter struct{}

var (
	mux = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.Get(uri, f)
}

func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.Post(uri, f)
}

func (*chiRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.Put(uri, f)
}

func (*chiRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mux.Delete(uri, f)
}

func (*chiRouter) SERVE(port string) {
	log.Default().Printf("Chi HTTP server running on port %v", port)
	http.ListenAndServe(port, mux)
}
