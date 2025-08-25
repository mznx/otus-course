package auth

import (
	"context"
	"dialog-service/domain/auth"
)

type CheckTokenData struct {
	Token string
}

type CheckTokenResult struct {
	UserID string
}

type UserCheckTokenService struct {
	AuthApi auth.Api
}

func NewUserCheckTokenService(authApi auth.Api) *UserCheckTokenService {
	return &UserCheckTokenService{AuthApi: authApi}
}

func (s *UserCheckTokenService) Handle(ctx context.Context, data *CheckTokenData) (*CheckTokenResult, error) {
	userId, err := s.AuthApi.ValidateToken(ctx, data.Token)

	if err != nil {
		return nil, err
	}

	return &CheckTokenResult{UserID: userId}, nil
}
