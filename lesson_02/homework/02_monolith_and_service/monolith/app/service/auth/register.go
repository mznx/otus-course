package auth

import (
	"context"
	"monolith/domain/user"
	"monolith/helper"
)

type RegisterData struct {
	FirstName  string
	SecondName string
	Password   string
}

type RegisterResult struct {
	UserID string
}

type UserRegisterService struct {
	userRepository user.Repository
}

func NewUserRegisterService(userRepository user.Repository) *UserRegisterService {
	return &UserRegisterService{userRepository: userRepository}
}

func (s *UserRegisterService) Handle(ctx context.Context, data *RegisterData) (*RegisterResult, error) {
	u := user.NewUser(data.FirstName, data.SecondName)

	passHash := helper.HashingPassword(data.Password)

	if err := s.userRepository.Create(ctx, u, passHash); err != nil {
		return nil, err
	}

	return &RegisterResult{UserID: u.ID}, nil
}
