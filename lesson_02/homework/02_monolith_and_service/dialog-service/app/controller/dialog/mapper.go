package dialog

import (
	"dialog-service/helper"
	dialog_service "dialog-service/service/dialog"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GenerateSendMessageData(r *http.Request) (*dialog_service.SendMessageData, error) {
	var body SendMessageRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &dialog_service.SendMessageData{
		SenderID:    helper.GetAuthorizedUserId(r),
		RecipientID: chi.URLParam(r, "user_id"),
		Text:        body.Text,
	}, nil
}

func GenerateGetListMessagesData(r *http.Request) (*dialog_service.GetListMessagesData, error) {
	return &dialog_service.GetListMessagesData{
		SenderID:    helper.GetAuthorizedUserId(r),
		RecipientID: chi.URLParam(r, "user_id"),
	}, nil
}

func GenerateGetListMessagesResponse(result *dialog_service.GetListMessagesResult, senderId string, recipientId string) []*GetListMessagesResponse {
	response := []*GetListMessagesResponse{}

	for _, message := range result.Messages {
		var from, to string

		if message.SenderID == senderId {
			from = senderId
			to = recipientId
		} else {
			from = recipientId
			to = senderId
		}

		response = append(response, &GetListMessagesResponse{
			From: from,
			To:   to,
			Text: message.Text,
		})
	}

	return response
}
