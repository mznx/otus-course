package service

import (
	"dialog-service/infrastructure/storage"
	"dialog-service/service/auth"
	"dialog-service/service/dialog"
)

type Service struct {
	Auth   *auth.AuthService
	Dialog *dialog.DialogService
}

func NewService(repositories *storage.Repository) *Service {
	return &Service{
		Auth:   auth.NewService(repositories),
		Dialog: dialog.NewService(repositories),
	}
}
