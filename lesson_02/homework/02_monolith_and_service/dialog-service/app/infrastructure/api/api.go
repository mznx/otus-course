package api

import (
	"dialog-service/domain/auth"
	auth_api "dialog-service/infrastructure/api/auth"
	"dialog-service/infrastructure/config"
)

type ExternalApi struct {
	Auth auth.Api
}

func NewExternalApi(config *config.Config) *ExternalApi {
	return &ExternalApi{
		Auth: auth_api.NewAuthApi(config),
	}
}
