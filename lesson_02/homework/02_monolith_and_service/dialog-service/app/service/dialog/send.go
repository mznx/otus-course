package dialog

import (
	"context"
	"dialog-service/domain/dialog"
	"dialog-service/domain/user"
	"errors"
)

type SendMessageData struct {
	SenderID    string
	RecipientID string
	Text        string
}

type SendMessageResult struct {
	MessageID string
}

type DialogSendMessageService struct {
	userApi          user.Api
	dialogRepository dialog.Repository
}

func NewDialogSendMessageService(dialogRepository dialog.Repository, userApi user.Api) *DialogSendMessageService {
	return &DialogSendMessageService{dialogRepository: dialogRepository, userApi: userApi}
}

func (s *DialogSendMessageService) Handle(ctx context.Context, data *SendMessageData) (*SendMessageResult, error) {
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

	return &SendMessageResult{MessageID: message.ID}, nil
}

func (s *DialogSendMessageService) checkIfRecipientExists(ctx context.Context, recipientId string) error {
	recipient, err := s.userApi.FindById(ctx, recipientId)

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
