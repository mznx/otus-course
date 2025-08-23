package dialog

import (
	"context"
	"errors"
	"monolith/domain/dialog"
	"monolith/domain/user"
)

type SendMessageRequest struct {
	SenderID    string `json:"sender_id"`
	RecipientID string `json:"recipient_id"`
	Text        string `json:"text"`
}

type SendMessageResponse struct {
	MessageID string `json:"message_id"`
}

type DialogSendMessageService struct {
	userRepository   user.Repository
	dialogRepository dialog.Repository
}

func NewDialogSendMessageService(userRepository user.Repository, dialogRepository dialog.Repository) *DialogSendMessageService {
	return &DialogSendMessageService{userRepository: userRepository, dialogRepository: dialogRepository}
}

func (s *DialogSendMessageService) Handle(ctx context.Context, data SendMessageRequest) (*SendMessageResponse, error) {
	if err := s.checkIfRecipientExists(ctx, data.RecipientID); err != nil {
		return nil, err
	}

	dialogId, err := s.getDialogId(ctx, data.SenderID, data.RecipientID)

	if err != nil {
		return nil, err
	}

	message := dialog.NewMessage(dialogId, data.SenderID, data.Text)

	if err := s.dialogRepository.SendMessage(ctx, message); err != nil {
		return nil, err
	}

	return &SendMessageResponse{MessageID: message.ID}, nil
}

func (s *DialogSendMessageService) checkIfRecipientExists(ctx context.Context, recipientId string) error {
	recipient, err := s.userRepository.FindById(ctx, recipientId)

	if err != nil {
		return err
	}

	if recipient == nil {
		return errors.New("recipient not found")
	}

	return nil
}

func (s *DialogSendMessageService) getDialogId(ctx context.Context, userId1 string, userId2 string) (string, error) {
	dialogId, err := s.dialogRepository.FindDialogBetweenUsers(ctx, userId1, userId2)

	if err != nil {
		return "", err
	}

	if dialogId == "" {
		dialogId, err = s.dialogRepository.CreatePrivateDialog(ctx, userId1, userId2)

		if err != nil {
			return "", err
		}
	}

	return dialogId, nil
}
