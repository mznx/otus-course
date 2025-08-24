package auth

import (
	"context"
)

type CheckTokenData struct {
	Token string
}

type CheckTokenResult struct {
	UserID string
}

type UserCheckTokenService struct{}

func NewUserCheckTokenService() *UserCheckTokenService {
	return &UserCheckTokenService{}
}

func (s *UserCheckTokenService) Handle(ctx context.Context, data *CheckTokenData) (*CheckTokenResult, error) {
	// TODO

	return &CheckTokenResult{UserID: ""}, nil
}
