package auth

import "monolith/infrastructure/storage"

type AuthService struct {
	UserLogin      *UserLoginService
	UserRegister   *UserRegisterService
	UserCheckToken *UserCheckTokenService
}

func NewService(repositories *storage.Repository) *AuthService {
	return &AuthService{
		UserLogin:      NewUserLoginService(repositories.User),
		UserRegister:   NewUserRegisterService(repositories.User),
		UserCheckToken: NewUserCheckTokenService(repositories.User),
	}
}
