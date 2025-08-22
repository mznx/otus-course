package service

import (
	"monolith/infrastructure/storage"
	"monolith/service/auth"
	"monolith/service/user"
)

type Service struct {
	Auth *auth.AuthService
	User *user.UserService
}

func NewService(repositories *storage.Repository) *Service {
	return &Service{
		Auth: auth.NewService(repositories),
		User: user.NewService(repositories),
	}
}
