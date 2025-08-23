package user

import (
	"encoding/json"
	"monolith/controller/middleware"
	"monolith/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRouter(router chi.Router, services *service.Service) {
	router.Get("/user/get/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		data, err := GenerateGetByIdData(r)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		res, err := services.User.GetById.Handle(r.Context(), data)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(GenerateGetByIdResponse(res))
	})

	router.Get("/user/search", func(w http.ResponseWriter, r *http.Request) {
		data, err := GenerateSearchData(r)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		res, err := services.User.Search.Handle(r.Context(), data)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(GenerateSearchResponse(res))
	})

	router.Route("/friend/set/{user_id}", func(r chi.Router) {
		r.Use(middleware.CheckToken(services))
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			data, err := GenerateAddFriendData(r)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			err = services.User.AddFriend.Handle(r.Context(), data)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err.Error())
				return
			}
		})
	})

	router.Route("/friend/delete/{user_id}", func(r chi.Router) {
		r.Use(middleware.CheckToken(services))
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			data, err := GenerateDeleteFriendData(r)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			err = services.User.DeleteFriend.Handle(r.Context(), data)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err.Error())
				return
			}
		})
	})
}
