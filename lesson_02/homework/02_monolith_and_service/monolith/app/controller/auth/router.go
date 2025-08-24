package auth

import (
	"encoding/json"
	"monolith/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AuthRouter(router chi.Router, services *service.Service) {
	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		data, err := GenerateLoginData(r)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		res, err := services.Auth.UserLogin.Handle(r.Context(), data)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(GenerateLoginResponse(res))
	})

	router.Post("/user/register", func(w http.ResponseWriter, r *http.Request) {
		data, err := GenerateRegisterData(r)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		res, err := services.Auth.UserRegister.Handle(r.Context(), data)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(GenerateRegisterResponse(res))
	})
}
