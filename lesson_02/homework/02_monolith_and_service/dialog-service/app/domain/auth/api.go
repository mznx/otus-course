package auth

import "context"

type Api interface {
	ValidateToken(ctx context.Context, token string) (string, error)
}
