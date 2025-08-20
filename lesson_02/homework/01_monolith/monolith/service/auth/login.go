package auth

import (
	"context"
	"monolith/domain/user"
	"monolith/helper"
)

type LoginRequest struct {
	UserID   string `json:"id"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserLoginService struct {
	userRepository user.Repository
}

func NewUserLoginService(userRepository user.Repository) *UserLoginService {
	return &UserLoginService{userRepository: userRepository}
}

func (s *UserLoginService) Handle(ctx context.Context, data LoginRequest) LoginResponse {
	passHash, err := s.userRepository.GetPasswordHash(ctx, data.UserID)

	if err != nil {
		// err
	}

	isValid := helper.IsValidPassword(passHash, data.Password)

	if !isValid {
		// err
	}

	return LoginResponse{Token: "user" + data.UserID}
}
