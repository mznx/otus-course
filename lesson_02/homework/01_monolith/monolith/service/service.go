package service

import (
	"monolith/infrastructure/storage"
	"monolith/service/auth"
)

type Service struct {
	Auth *auth.AuthService
}

func NewService(repositories *storage.Repository) *Service {
	return &Service{
		Auth: auth.NewService(repositories),
	}
}
