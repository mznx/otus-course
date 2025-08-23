package controller

import (
	"monolith/controller/auth"
	"monolith/controller/dialog"
	"monolith/controller/user"
	"monolith/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(services *service.Service) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	auth.AuthRouter(router, services)
	user.UserRouter(router, services)
	dialog.DialogRouter(router, services)

	return router
}
