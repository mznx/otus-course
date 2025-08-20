package controller

import (
	"monolith/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRouter(router chi.Router, services *service.Service) {
	router.Get("/user/get/{userId}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))

		// userId := chi.URLParam(r, "userId")

		// res := service.GetUserById(userId)

		// w.Write([]byte(res))
	})

	router.Get("/user/search", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	router.Put("/friend/set/{userId}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	router.Put("/friend/delete/{userId}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
}
