package dialog

import "monolith/infrastructure/storage"

type DialogService struct {
	SendMessage     *DialogSendMessageService
	GetListMessages *DialogGetListMessagesService
}

func NewService(repositories *storage.Repository) *DialogService {
	return &DialogService{
		SendMessage:     NewDialogSendMessageService(repositories.User, repositories.Dialog),
		GetListMessages: NewDialogGetListMessagesService(repositories.Dialog),
	}
}
