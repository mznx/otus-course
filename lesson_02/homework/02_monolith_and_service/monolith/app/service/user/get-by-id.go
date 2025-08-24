package user

import (
	"context"
	"monolith/domain/user"
)

type GetByIdData struct {
	UserID string
}

type GetByIdResult struct {
	User *user.User
}

type UserGetByIdService struct {
	userRepository user.Repository
}

func NewUserGetByIdService(userRepository user.Repository) *UserGetByIdService {
	return &UserGetByIdService{userRepository: userRepository}
}

func (s *UserGetByIdService) Handle(ctx context.Context, data *GetByIdData) (*GetByIdResult, error) {
	u, err := s.userRepository.FindById(ctx, data.UserID)

	if err != nil {
		return nil, err
	}

	return &GetByIdResult{User: u}, nil
}
