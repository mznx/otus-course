package dialog

import "github.com/google/uuid"

type Message struct {
	ID       string
	DialogID string
	SenderID string
	Text     string
}

func NewMessage(dialogId string, senderId string, text string) *Message {
	return &Message{
		ID:       uuid.New().String(),
		DialogID: dialogId,
		SenderID: senderId,
		Text:     text,
	}
}
