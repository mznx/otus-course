package user

import (
	"context"
	"database/sql"
	"errors"
	"monolith/domain/user"
	user_mapper "monolith/infrastructure/mapper/user"
	user_model "monolith/infrastructure/model/user"
	user_auth_model "monolith/infrastructure/model/user_auth"

	"github.com/jmoiron/sqlx"
)

type UserPgRepository struct {
	db *sqlx.DB
}

func NewUserPgRepository(db *sqlx.DB) *UserPgRepository {
	return &UserPgRepository{
		db: db,
	}
}

func (r *UserPgRepository) FindById(ctx context.Context, userId string) (*user.User, error) {
	u := user_model.User{}

	if err := r.db.Get(&u, "SELECT * FROM users WHERE id=$1", userId); err != nil {
		return nil, err
	}

	return user_mapper.ModelToUser(u), nil
}

func (r *UserPgRepository) GetPasswordHash(ctx context.Context, userId string) (string, error) {
	ua := user_auth_model.UserAuth{}

	if err := r.db.Get(&ua, "SELECT * FROM user_auth WHERE user_id=$1", userId); err != nil {
		return "", err
	}

	return ua.PassHash, nil
}

func (r *UserPgRepository) Create(ctx context.Context, user *user.User, passHash string) error {
	return runInTx(r.db, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, "INSERT INTO users (id, first_name, second_name) VALUES ($1, $2, $3)", user.ID, user.FirstName, user.SecondName)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, "INSERT INTO user_auth (user_id, pass_hash) VALUES ($1, $2)", user.ID, passHash)
		if err != nil {
			return err
		}

		return nil
	})
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
