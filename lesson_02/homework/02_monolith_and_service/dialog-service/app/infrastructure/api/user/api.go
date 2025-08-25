package user

import (
	"context"
	"dialog-service/domain/user"
	"dialog-service/infrastructure/config"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type UserApi struct {
	baseUrl string
}

func NewUserApi(config *config.Config) *UserApi {
	return &UserApi{
		baseUrl: config.Services.User,
	}
}

func (a *UserApi) FindById(ctx context.Context, userId string) (*user.User, error) {
	res, err := a.doGetRequest(ctx, fmt.Sprintf("/user/get/%s", userId))
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	var result FindByIdResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return user.NewUser(result.UserID), nil
}

func (a *UserApi) doGetRequest(ctx context.Context, route string) (*http.Response, error) {
	url := a.baseUrl + route

	client := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, err
}
