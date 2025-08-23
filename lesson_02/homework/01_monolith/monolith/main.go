package main

import (
	"monolith/controller"
	"monolith/infrastructure/storage"
	"monolith/infrastructure/storage/postgres"
	"monolith/service"
	"net/http"
)

func main() {
	s := storage.Connect()
	defer s.Disconnect()

	repositories := postgres.NewRepository(s.GetDB())

	services := service.NewService(repositories)

	router := controller.NewRouter(services)

	http.ListenAndServe(":3000", router)
}
