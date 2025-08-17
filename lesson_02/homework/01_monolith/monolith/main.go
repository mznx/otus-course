package main

import (
	"monolith/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := initRouter()

	http.ListenAndServe(":3000", r)
}

func initRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	controller.UserRouter(router)

	return router
}
