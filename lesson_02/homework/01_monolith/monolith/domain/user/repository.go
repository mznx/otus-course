package user

import "context"

type Repository interface {
	FindById(ctx context.Context, userId string) (*User, error)

	FindByName(ctx context.Context, firstName string, lastName string) ([]*User, error)

	GetPasswordHash(ctx context.Context, userId string) (string, error)

	CheckIfUsersAreFriends(ctx context.Context, userId string, friendId string) (bool, error)

	UpdateAuthToken(ctx context.Context, userId string, token string) error

	AddFriend(ctx context.Context, userId string, friendId string) error

	DeleteFriend(ctx context.Context, userId string, friendId string) error

	Create(ctx context.Context, user *User, passHash string) error
}
