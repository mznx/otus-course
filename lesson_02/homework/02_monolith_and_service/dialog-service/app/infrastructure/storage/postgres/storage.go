package postgres

import (
	"dialog-service/infrastructure/repository/dialog"
	"dialog-service/infrastructure/storage"

	"github.com/jmoiron/sqlx"
)

func NewRepository(db *sqlx.DB) *storage.Repository {
	return &storage.Repository{
		Dialog: dialog.NewDialogPgRepository(db),
	}
}
