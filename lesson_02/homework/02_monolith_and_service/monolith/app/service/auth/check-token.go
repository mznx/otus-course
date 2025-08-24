package auth

import (
	"context"
	"monolith/domain/user"
)

type CheckTokenData struct {
	Token string
}

type CheckTokenResult struct {
	UserID string
}

type UserCheckTokenService struct {
	userRepository user.Repository
}

func NewUserCheckTokenService(userRepository user.Repository) *UserCheckTokenService {
	return &UserCheckTokenService{userRepository: userRepository}
}

func (s *UserCheckTokenService) Handle(ctx context.Context, data *CheckTokenData) (*CheckTokenResult, error) {
	u, err := s.userRepository.FindByToken(ctx, data.Token)

	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, nil
	}

	return &CheckTokenResult{UserID: u.ID}, nil
}
