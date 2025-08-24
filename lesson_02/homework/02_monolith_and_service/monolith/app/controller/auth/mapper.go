package auth

import (
	"encoding/json"
	"monolith/helper"
	auth_service "monolith/service/auth"
	"net/http"
)

func GenerateLoginData(r *http.Request) (*auth_service.LoginData, error) {
	var body LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &auth_service.LoginData{
		UserID:   body.UserID,
		Password: body.Password,
	}, nil
}

func GenerateLoginResponse(result *auth_service.LoginResult) *LoginResponse {
	return &LoginResponse{
		Token: result.Token,
	}
}

func GenerateRegisterData(r *http.Request) (*auth_service.RegisterData, error) {
	var body RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &auth_service.RegisterData{
		FirstName:  body.FirstName,
		SecondName: body.SecondName,
		Password:   body.Password,
	}, nil
}

func GenerateRegisterResponse(result *auth_service.RegisterResult) *RegisterResponse {
	return &RegisterResponse{
		UserID: result.UserID,
	}
}

func GenerateTokenValidateResponse(r *http.Request) *TokenValidateResponse {
	return &TokenValidateResponse{
		UserID: helper.GetAuthorizedUserId(r),
	}
}
