package dialog

import "monolith/domain/dialog"

func GenerateGetListMessagesResponse(messages []*dialog.Message, senderId string, recipientId string) []*GetListMessagesResponse {
	result := []*GetListMessagesResponse{}

	for _, message := range messages {
		var from, to string

		if message.SenderID == senderId {
			from = senderId
			to = recipientId
		} else {
			from = recipientId
			to = senderId
		}

		result = append(result, &GetListMessagesResponse{
			From: from,
			To:   to,
			Text: message.Text,
		})
	}

	return result
}
