package auth

import (
	"context"
	"errors"
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

func (s *UserLoginService) Handle(ctx context.Context, data LoginRequest) (*LoginResponse, error) {
	passHash, err := s.userRepository.GetPasswordHash(ctx, data.UserID)

	if err != nil {
		return nil, err
	}

	isValid := helper.IsValidPassword(passHash, data.Password)

	if !isValid {
		return nil, errors.New("incorrect password")
	}

	token := helper.GenerateAuthToken(data.UserID)

	err = s.userRepository.UpdateAuthToken(ctx, data.UserID, token)

	if err != nil {
		return nil, err
	}

	return &LoginResponse{Token: token}, nil
}
