package user

import (
	"context"
	"monolith/domain/user"
)

type GetByIdRequest struct {
	UserID string `json:"id"`
}

type GetByIdResponse struct {
	UserID     string `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

type UserGetByIdService struct {
	userRepository user.Repository
}

func NewUserGetByIdService(userRepository user.Repository) *UserGetByIdService {
	return &UserGetByIdService{userRepository: userRepository}
}

func (s *UserGetByIdService) Handle(ctx context.Context, data GetByIdRequest) (*GetByIdResponse, error) {
	u, err := s.userRepository.FindById(ctx, data.UserID)

	if err != nil {
		return nil, err
	}

	return &GetByIdResponse{UserID: u.ID, FirstName: u.FirstName, SecondName: u.SecondName}, nil
}
