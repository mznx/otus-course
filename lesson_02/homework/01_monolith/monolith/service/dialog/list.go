package dialog

import (
	"context"
	"errors"
	"monolith/domain/dialog"
)

type GetListMessagesRequest struct {
	SenderID    string `json:"sender_id"`
	RecipientID string `json:"recipient_id"`
}

type GetListMessagesResponse struct {
	Messages []*dialog.Message
}

type DialogGetListMessagesService struct {
	dialogRepository dialog.Repository
}

func NewDialogGetListMessagesService(dialogRepository dialog.Repository) *DialogGetListMessagesService {
	return &DialogGetListMessagesService{dialogRepository: dialogRepository}
}

func (s *DialogGetListMessagesService) Handle(ctx context.Context, data GetListMessagesRequest) (*GetListMessagesResponse, error) {
	dialogId, err := s.dialogRepository.FindDialogBetweenUsers(ctx, data.SenderID, data.RecipientID)

	if err != nil {
		return nil, err
	}

	if dialogId == "" {
		return nil, errors.New("dialog not found")
	}

	messages, err := s.dialogRepository.GetDialogMessages(ctx, dialogId)

	if err != nil {
		return nil, err
	}

	return &GetListMessagesResponse{Messages: messages}, nil
}
