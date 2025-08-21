package auth

import (
	"context"
	"monolith/domain/user"
	"monolith/helper"
)

type RegisterRequest struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Password   string `json:"password"`
}

type RegisterResponse struct {
	UserID string `json:"user_id"`
}

type UserRegisterService struct {
	userRepository user.Repository
}

func NewUserRegisterService(userRepository user.Repository) *UserRegisterService {
	return &UserRegisterService{userRepository: userRepository}
}

func (s *UserRegisterService) Handle(ctx context.Context, data RegisterRequest) (*RegisterResponse, error) {
	u := user.NewUser(data.FirstName, data.SecondName)

	passHash := helper.HashingPassword(data.Password)

	if err := s.userRepository.Create(ctx, u, passHash); err != nil {
		return nil, err
	}

	return &RegisterResponse{UserID: u.ID}, nil
}
