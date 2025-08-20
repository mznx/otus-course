package auth

import (
	"context"
	"monolith/domain/user"
)

type RegisterRequest struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Password   string `json:"password"`
}

type RegisterResponse struct {
	UserId string `json:"user_id"`
}

type UserRegisterService struct {
	userRepository user.Repository
}

func NewUserRegisterService(userRepository user.Repository) *UserRegisterService {
	return &UserRegisterService{userRepository: userRepository}
}

func (s *UserRegisterService) Handle(ctx context.Context, data RegisterRequest) RegisterResponse {
	u := user.NewUser(data.FirstName, data.SecondName)

	if err := s.userRepository.Create(ctx, u, data.Password); err != nil {
		// err
	}

	return RegisterResponse{UserId: u.ID}
}
