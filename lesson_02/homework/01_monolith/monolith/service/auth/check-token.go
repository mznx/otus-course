package auth

import (
	"context"
	"monolith/domain/user"
)

type CheckTokenRequest struct {
	Token string `json:"token"`
}

type CheckTokenResponse struct {
	UserID string `json:"user_id"`
}

type UserCheckTokenService struct {
	userRepository user.Repository
}

func NewUserCheckTokenService(userRepository user.Repository) *UserCheckTokenService {
	return &UserCheckTokenService{userRepository: userRepository}
}

func (s *UserCheckTokenService) Handle(ctx context.Context, data CheckTokenRequest) (*CheckTokenResponse, error) {
	u, err := s.userRepository.FindByToken(ctx, data.Token)

	if err != nil {
		return nil, err
	}

	return &CheckTokenResponse{UserID: u.ID}, nil
}
