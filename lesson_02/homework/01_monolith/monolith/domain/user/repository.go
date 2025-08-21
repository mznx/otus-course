package user

import "context"

type Repository interface {
	FindById(ctx context.Context, userId string) (*User, error)

	GetPasswordHash(ctx context.Context, userId string) (string, error)

	UpdateAuthToken(ctx context.Context, userId string, token string) error

	Create(ctx context.Context, user *User, passHash string) error
}
