package user

import (
	"context"
	"monolith/domain/user"
)

type SearchData struct {
	FirstName  string
	SecondName string
}

type SearchResult struct {
	Users []*user.User
}

type UserSearchService struct {
	userRepository user.Repository
}

func NewUserSearchService(userRepository user.Repository) *UserSearchService {
	return &UserSearchService{userRepository: userRepository}
}

func (s *UserSearchService) Handle(ctx context.Context, data *SearchData) (*SearchResult, error) {
	users, err := s.userRepository.FindByName(ctx, data.FirstName, data.SecondName)

	if err != nil {
		return nil, err
	}

	return &SearchResult{Users: users}, nil
}
