package controller

import (
	"monolith/controller/auth"
	"monolith/controller/dialog"
	"monolith/controller/post"
	"monolith/controller/user"
	"monolith/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(services *service.Service) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	auth.AuthRouter(router, services)
	user.UserRouter(router, services)
	post.PostRouter(router, services)
	dialog.DialogRouter(router, services)

	return router
}
