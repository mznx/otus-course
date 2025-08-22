package controller

import (
	"encoding/json"
	"monolith/service"
	"monolith/service/user"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRouter(router chi.Router, services *service.Service) {
	router.Get("/user/get/{userId}", func(w http.ResponseWriter, r *http.Request) {
		data := user.GetByIdRequest{UserID: chi.URLParam(r, "userId")}

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

	router.Put("/friend/set/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userId := ""

		data := user.AddFriendRequest{
			UserID:   userId,
			FriendID: r.URL.Query().Get("user_id"),
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

	router.Put("/friend/delete/{userId}", func(w http.ResponseWriter, r *http.Request) {
		userId := ""

		data := user.DeleteFriendRequest{
			UserID:   userId,
			FriendID: r.URL.Query().Get("user_id"),
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
}
