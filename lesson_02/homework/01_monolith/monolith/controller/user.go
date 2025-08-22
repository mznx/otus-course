package controller

import (
	"encoding/json"
	"monolith/controller/middleware"
	"monolith/helper"
	"monolith/service"
	"monolith/service/user"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRouter(router chi.Router, services *service.Service) {
	router.Get("/user/get/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		data := user.GetByIdRequest{UserID: chi.URLParam(r, "user_id")}

		ctx := r.Context()

		res, err := services.User.GetById.Handle(ctx, data)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	router.Get("/user/search", func(w http.ResponseWriter, r *http.Request) {
		data := user.SearchRequest{
			FirstName:  r.URL.Query().Get("first_name"),
			SecondName: r.URL.Query().Get("last_name"),
		}

		ctx := r.Context()

		res, err := services.User.Search.Handle(ctx, data)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	router.Route("/friend/set/{user_id}", func(r chi.Router) {
		r.Use(middleware.CheckToken(services))
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			data := user.AddFriendRequest{
				UserID:   helper.GetAuthorizedUserId(r),
				FriendID: chi.URLParam(r, "user_id"),
			}

			ctx := r.Context()

			err := services.User.AddFriend.Handle(ctx, data)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err.Error())
				return
			}

			w.Header().Set("Content-Type", "application/json")
		})
	})

	router.Route("/friend/delete/{user_id}", func(r chi.Router) {
		r.Use(middleware.CheckToken(services))
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			data := user.DeleteFriendRequest{
				UserID:   helper.GetAuthorizedUserId(r),
				FriendID: chi.URLParam(r, "user_id"),
			}

			ctx := r.Context()

			err := services.User.DeleteFriend.Handle(ctx, data)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err.Error())
				return
			}

			w.Header().Set("Content-Type", "application/json")
		})
	})
}
