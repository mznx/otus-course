package service

import (
	"dialog-service/infrastructure/api"
	"dialog-service/infrastructure/storage"
	"dialog-service/service/auth"
	"dialog-service/service/dialog"
)

type Service struct {
	Auth   *auth.AuthService
	Dialog *dialog.DialogService
}

func NewService(repositories *storage.Repository, api *api.ExternalApi) *Service {
	return &Service{
		Auth:   auth.NewService(api),
		Dialog: dialog.NewService(repositories),
	}
}
