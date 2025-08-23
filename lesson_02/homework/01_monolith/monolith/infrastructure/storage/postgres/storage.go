package postgres

import (
	"monolith/infrastructure/repository/dialog"
	"monolith/infrastructure/repository/user"
	"monolith/infrastructure/storage"

	"github.com/jmoiron/sqlx"
)

func NewRepository(db *sqlx.DB) *storage.Repository {
	return &storage.Repository{
		User:   user.NewUserPgRepository(db),
		Dialog: dialog.NewDialogPgRepository(db),
	}
}
