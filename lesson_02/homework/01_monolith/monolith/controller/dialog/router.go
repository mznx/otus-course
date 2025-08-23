package dialog

import (
	"encoding/json"
	"monolith/controller/middleware"
	"monolith/helper"
	"monolith/service"
	"monolith/service/dialog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func DialogRouter(router chi.Router, services *service.Service) {
	router.Route("/dialog/{user_id}/send", func(r chi.Router) {
		r.Use(middleware.CheckToken(services))
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			var body struct {
				Text string `json:"text"`
			}

			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			data := dialog.SendMessageRequest{
				SenderID:    helper.GetAuthorizedUserId(r),
				RecipientID: chi.URLParam(r, "user_id"),
				Text:        body.Text,
			}

			ctx := r.Context()

			_, err := services.Dialog.SendMessage.Handle(ctx, data)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err.Error())
				return
			}

			w.Header().Set("Content-Type", "application/json")
		})
	})

	router.Route("/dialog/{user_id}/list", func(r chi.Router) {
		r.Use(middleware.CheckToken(services))
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			senderId := helper.GetAuthorizedUserId(r)
			recipientId := chi.URLParam(r, "user_id")

			data := dialog.GetListMessagesRequest{
				SenderID:    senderId,
				RecipientID: recipientId,
			}

			ctx := r.Context()

			res, err := services.Dialog.GetListMessages.Handle(ctx, data)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err.Error())
				return
			}

			resp := GenerateGetListMessagesResponse(res.Messages, senderId, recipientId)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		})
	})
}
