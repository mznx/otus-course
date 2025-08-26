package postgres

import (
	"monolith/infrastructure/repository/post"
	"monolith/infrastructure/repository/user"
	"monolith/infrastructure/storage"

	"github.com/jmoiron/sqlx"
)

func NewRepository(db *sqlx.DB) *storage.Repository {
	return &storage.Repository{
		User: user.NewUserPgRepository(db),
		Post: post.NewPostPgRepository(db),
	}
}
