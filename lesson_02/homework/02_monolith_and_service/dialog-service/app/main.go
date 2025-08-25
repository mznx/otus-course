package main

import (
	"dialog-service/controller"
	"dialog-service/infrastructure/api"
	"dialog-service/infrastructure/config"
	"dialog-service/infrastructure/storage"
	"dialog-service/infrastructure/storage/postgres"
	"dialog-service/service"
	"net/http"
)

func main() {
	config := config.ReadConfig()

	s := storage.Connect(config)
	defer s.Disconnect()

	repositories := postgres.NewRepository(s.GetDB())

	externalApi := api.NewExternalApi(config)

	services := service.NewService(repositories, externalApi)

	router := controller.NewRouter(services)

	http.ListenAndServe(":3000", router)
}
