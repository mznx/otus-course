package auth

import (
	"context"
	"dialog-service/infrastructure/config"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type AuthApi struct {
	baseUrl string
}

func NewAuthApi(config *config.Config) *AuthApi {
	return &AuthApi{
		baseUrl: config.Services.Auth,
	}
}

func (a *AuthApi) ValidateToken(ctx context.Context, token string) (string, error) {
	res, err := a.doRequest(ctx, "/user/token/validate", token)
	if err != nil {
		return "", err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	var result ValidateTokenResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.UserID, nil

}

func (a *AuthApi) doRequest(ctx context.Context, route string, token string) (*http.Response, error) {
	url := a.baseUrl + route

	client := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, err
}
