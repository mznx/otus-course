package main

import (
	"monolith/controller"
	"monolith/infrastructure/storage"
	"monolith/infrastructure/storage/postgres"
	"monolith/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	s := storage.Connect()
	defer s.Disconnect()

	repositories := postgres.NewRepository(s.GetDB())

	services := service.NewService(repositories)

	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	controller.AuthRouter(router, services)
	controller.UserRouter(router, services)

	http.ListenAndServe(":3000", router)
}
