package dialog

import (
	"dialog-service/infrastructure/api"
	"dialog-service/infrastructure/storage"
)

type DialogService struct {
	SendMessage     *DialogSendMessageService
	GetListMessages *DialogGetListMessagesService
}

func NewService(repositories *storage.Repository, api *api.ExternalApi) *DialogService {
	return &DialogService{
		SendMessage:     NewDialogSendMessageService(repositories.Dialog, api.User),
		GetListMessages: NewDialogGetListMessagesService(repositories.Dialog),
	}
}
