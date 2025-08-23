package service

import (
	"monolith/infrastructure/storage"
	"monolith/service/auth"
	"monolith/service/dialog"
	"monolith/service/user"
)

type Service struct {
	Auth   *auth.AuthService
	User   *user.UserService
	Dialog *dialog.DialogService
}

func NewService(repositories *storage.Repository) *Service {
	return &Service{
		Auth:   auth.NewService(repositories),
		User:   user.NewService(repositories),
		Dialog: dialog.NewService(repositories),
	}
}
