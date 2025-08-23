package post

import (
	"encoding/json"
	"monolith/controller/middleware"
	"monolith/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func PostRouter(router chi.Router, services *service.Service) {
	router.Route("/post/create", func(r chi.Router) {
		r.Use(middleware.CheckToken(services))
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			data, err := GenerateCreatePostData(r)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			result, err := services.Post.CreatePost.Handle(r.Context(), data)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err.Error())
				return
			}

			w.Write([]byte(result.PostID))
		})
	})

	router.Route("/post/update", func(r chi.Router) {
		r.Use(middleware.CheckToken(services))
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			data, err := GenerateUpdatePostData(r)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if err := services.Post.UpdatePost.Handle(r.Context(), data); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err.Error())
				return
			}
		})
	})

	router.Route("/post/delete/{post_id}", func(r chi.Router) {
		r.Use(middleware.CheckToken(services))
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			data, err := GenerateDeletePostData(r)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if err := services.Post.DeletePost.Handle(r.Context(), data); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err.Error())
				return
			}
		})
	})

	router.Get("/post/get/{post_id}", func(w http.ResponseWriter, r *http.Request) {
		data, err := GenerateGetByIdtData(r)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		res, err := services.Post.GetById.Handle(r.Context(), data)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(GenerateGetByIdtResponse(res))

	})

	router.Route("/post/feed", func(r chi.Router) {
		r.Use(middleware.CheckToken(services))
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			data, err := GenerateGetFeedData(r)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			res, err := services.Post.GetFeed.Handle(r.Context(), data)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err.Error())
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(GenerateGetFeedResponse(res))
		})
	})
}
