package user

import "context"

type Api interface {
	FindById(ctx context.Context, userId string) (*User, error)
}
