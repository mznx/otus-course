package user

import (
	"context"
	"database/sql"
	"errors"
	"monolith/domain/user"
	user_mapper "monolith/infrastructure/mapper/user"
	friend_model "monolith/infrastructure/model/friend"
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
	var user user_model.User

	if err := r.db.Get(&user, "SELECT * FROM users WHERE id=$1", userId); err != nil {
		return nil, err
	}

	return user_mapper.ModelToUser(user), nil
}

func (r *UserPgRepository) FindByName(ctx context.Context, firstName string, secondName string) ([]*user.User, error) {
	var users []user_model.User

	if err := r.db.SelectContext(ctx, &users, "SELECT * FROM users WHERE first_name=$1 AND second_name=$2", firstName, secondName); err != nil {
		return nil, err
	}

	return user_mapper.ModelsToUsers(users), nil
}

func (r *UserPgRepository) FindByToken(ctx context.Context, token string) (*user.User, error) {
	var user user_model.User

	if err := r.db.Get(&user, "SELECT users.* FROM users INNER JOIN user_auth ON users.id = user_auth.user_id WHERE user_auth.token=$1", token); err != nil {
		return nil, err
	}

	return user_mapper.ModelToUser(user), nil
}

func (r *UserPgRepository) GetPasswordHash(ctx context.Context, userId string) (string, error) {
	var userAuth user_auth_model.UserAuth

	if err := r.db.Get(&userAuth, "SELECT * FROM user_auth WHERE user_id=$1", userId); err != nil {
		return "", err
	}

	return userAuth.PassHash, nil
}

func (r *UserPgRepository) CheckIfUsersAreFriends(ctx context.Context, userId string, friendId string) (bool, error) {
	var friend friend_model.Friend

	err := r.db.Get(&friend, "SELECT * FROM friends WHERE user_id=$1 AND friend_id=$2", userId, friendId)

	switch err {
	case nil:
		return true, nil
	case sql.ErrNoRows:
		return false, nil
	default:
		return false, err
	}
}

func (r *UserPgRepository) UpdateAuthToken(ctx context.Context, userId string, token string) error {
	if _, err := r.db.ExecContext(ctx, "UPDATE user_auth SET token=$1 WHERE user_id=$2", token, userId); err != nil {
		return err
	}

	return nil
}

func (r *UserPgRepository) AddFriend(ctx context.Context, userId string, friendId string) error {
	return runInTx(r.db, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, "INSERT INTO friends (user_id, friend_id) VALUES ($1, $2)", userId, friendId)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, "INSERT INTO friends (user_id, friend_id) VALUES ($1, $2)", friendId, userId)
		if err != nil {
			return err
		}

		return nil
	})
}

func (r *UserPgRepository) DeleteFriend(ctx context.Context, userId string, friendId string) error {
	return runInTx(r.db, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, "DELETE FROM friends WHERE user_id=$1 AND friend_id=$2", userId, friendId)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, "DELETE FROM friends WHERE user_id=$1 AND friend_id=$2", friendId, userId)
		if err != nil {
			return err
		}

		return nil
	})
}

func (r *UserPgRepository) Create(ctx context.Context, user *user.User, passHash string) error {
	return runInTx(r.db, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, "INSERT INTO users (id, first_name, second_name) VALUES ($1, $2, $3)", user.ID, user.FirstName, user.SecondName)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, "INSERT INTO user_auth (user_id, pass_hash, token) VALUES ($1, $2, $3)", user.ID, passHash, passHash)
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
