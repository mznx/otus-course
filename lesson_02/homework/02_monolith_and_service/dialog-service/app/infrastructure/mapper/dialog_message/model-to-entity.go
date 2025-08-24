package dialog_message_mapper

import (
	"dialog-service/domain/dialog"
	"dialog-service/infrastructure/model/dialog_message"
)

func ModelToMessage(m dialog_message.DialogMessage) *dialog.Message {
	return &dialog.Message{
		ID:       m.ID,
		DialogID: m.DialogID,
		SenderID: m.SenderID,
		Text:     m.Message,
	}
}

func ModelsToMessages(messages []dialog_message.DialogMessage) []*dialog.Message {
	result := []*dialog.Message{}

	for _, m := range messages {
		result = append(result, ModelToMessage(m))
	}

	return result
}
