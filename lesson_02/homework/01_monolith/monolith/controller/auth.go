package controller

import (
	"encoding/json"
	"monolith/service"
	"monolith/service/auth"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AuthRouter(router chi.Router, services *service.Service) {
	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		body := auth.LoginRequest{}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		res, err := services.Auth.UserLogin.Handle(ctx, body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	router.Post("/user/register", func(w http.ResponseWriter, r *http.Request) {
		body := auth.RegisterRequest{}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		res, err := services.Auth.UserRegister.Handle(ctx, body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})
}
