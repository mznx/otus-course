package api

import (
	"dialog-service/domain/auth"
	"dialog-service/domain/user"
	auth_api "dialog-service/infrastructure/api/auth"
	user_api "dialog-service/infrastructure/api/user"
	"dialog-service/infrastructure/config"
)

type ExternalApi struct {
	Auth auth.Api
	User user.Api
}

func NewExternalApi(config *config.Config) *ExternalApi {
	return &ExternalApi{
		Auth: auth_api.NewAuthApi(config),
		User: user_api.NewUserApi(config),
	}
}
