package dialog

import (
	"dialog-service/controller/middleware"
	"dialog-service/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func DialogRouter(router chi.Router, services *service.Service) {
	router.Route("/dialog/{user_id}/send", func(r chi.Router) {
		r.Use(middleware.CheckToken(services))
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			data, err := GenerateSendMessageData(r)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			_, err = services.Dialog.SendMessage.Handle(r.Context(), data)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err.Error())
				return
			}
		})
	})

	router.Route("/dialog/{user_id}/list", func(r chi.Router) {
		r.Use(middleware.CheckToken(services))
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			data, err := GenerateGetListMessagesData(r)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			res, err := services.Dialog.GetListMessages.Handle(r.Context(), data)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err.Error())
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(GenerateGetListMessagesResponse(res, data.SenderID, data.RecipientID))
		})
	})
}
