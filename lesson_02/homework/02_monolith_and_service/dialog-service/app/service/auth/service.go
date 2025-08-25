package auth

import "dialog-service/infrastructure/api"

type AuthService struct {
	UserCheckToken *UserCheckTokenService
}

func NewService(api *api.ExternalApi) *AuthService {
	return &AuthService{
		UserCheckToken: NewUserCheckTokenService(api.Auth),
	}
}
