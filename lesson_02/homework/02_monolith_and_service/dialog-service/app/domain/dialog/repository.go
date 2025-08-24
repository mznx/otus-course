package dialog

import "context"

type Repository interface {
	FindDialogBetweenUsers(ctx context.Context, userId1 string, userId2 string) (string, error)

	GetDialogMessages(ctx context.Context, dialogId string) ([]*Message, error)

	CreatePrivateDialog(ctx context.Context, userId1 string, userId2 string) (string, error)

	SendMessage(ctx context.Context, message *Message) error
}
