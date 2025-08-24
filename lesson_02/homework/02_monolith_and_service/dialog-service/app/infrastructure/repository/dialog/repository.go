package dialog

import (
	"context"
	"database/sql"
	"dialog-service/domain/dialog"
	"dialog-service/helper"
	dialog_message_mapper "dialog-service/infrastructure/mapper/dialog_message"
	"dialog-service/infrastructure/model/dialog_message"
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type DialogPgRepository struct {
	db *sqlx.DB
}

func NewDialogPgRepository(db *sqlx.DB) *DialogPgRepository {
	return &DialogPgRepository{
		db: db,
	}
}

func (r *DialogPgRepository) FindDialogBetweenUsers(ctx context.Context, userId1 string, userId2 string) (string, error) {
	userPairHash := helper.GenerateUserPairHash(userId1, userId2)

	var dialogId string

	err := r.db.GetContext(ctx, &dialogId, "SELECT dialog_id FROM dialog_private WHERE user_pair_hash=$1", userPairHash)

	switch err {
	case nil:
		return dialogId, nil
	case sql.ErrNoRows:
		return "", nil
	default:
		return "", err
	}
}

func (r *DialogPgRepository) GetDialogMessages(ctx context.Context, dialogId string) ([]*dialog.Message, error) {
	var messages []dialog_message.DialogMessage

	if err := r.db.SelectContext(ctx, &messages, "SELECT * FROM dialog_messages WHERE dialog_id=$1", dialogId); err != nil {
		return nil, err
	}

	return dialog_message_mapper.ModelsToMessages(messages), nil
}

func (r *DialogPgRepository) CreatePrivateDialog(ctx context.Context, userId1 string, userId2 string) (string, error) {
	dialogId := uuid.New().String()

	userPairHash := helper.GenerateUserPairHash(userId1, userId2)

	err := runInTx(r.db, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, `INSERT INTO dialogs (id) VALUES ($1)`, dialogId)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, `INSERT INTO dialog_members (user_id, dialog_id) VALUES ($1, $2)`, userId1, dialogId)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, `INSERT INTO dialog_members (user_id, dialog_id) VALUES ($1, $2)`, userId2, dialogId)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, `
			INSERT INTO dialog_private (user_pair_hash, dialog_id, user_id_1, user_id_2)
			VALUES ($1, $2, $3, $4)`, userPairHash, dialogId, userId1, userId2)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return dialogId, nil
}

func (r *DialogPgRepository) SendMessage(ctx context.Context, message *dialog.Message) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO dialog_messages (id, dialog_id, sender_id, message)
		VALUES ($1, $2, $3, $4)`, message.ID, message.DialogID, message.SenderID, message.Text)
	if err != nil {
		return err
	}

	return nil
}

func runInTx(db *sqlx.DB, fn func(tx *sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	err = fn(tx)
	if err == nil {
		return tx.Commit()
	}

	rollbackErr := tx.Rollback()
	if rollbackErr != nil {
		return errors.Join(err, rollbackErr)
	}

	return err
}
