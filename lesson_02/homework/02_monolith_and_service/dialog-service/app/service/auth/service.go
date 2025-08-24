package auth

import "dialog-service/infrastructure/storage"

type AuthService struct {
	UserCheckToken *UserCheckTokenService
}

func NewService(repositories *storage.Repository) *AuthService {
	return &AuthService{
		UserCheckToken: NewUserCheckTokenService(),
	}
}
