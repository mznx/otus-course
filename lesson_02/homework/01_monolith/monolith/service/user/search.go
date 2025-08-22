package user

import (
	"context"
	"monolith/domain/user"
)

type SearchRequest struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"last_name"`
}

type SearchResponse struct {
	UserID     string `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

type UserSearchService struct {
	userRepository user.Repository
}

func NewUserSearchService(userRepository user.Repository) *UserSearchService {
	return &UserSearchService{userRepository: userRepository}
}

func (s *UserSearchService) Handle(ctx context.Context, data SearchRequest) ([]*SearchResponse, error) {
	users, err := s.userRepository.FindByName(ctx, data.FirstName, data.SecondName)

	if err != nil {
		return nil, err
	}

	result := []*SearchResponse{}

	for _, u := range users {
		result = append(result, &SearchResponse{UserID: u.ID, FirstName: u.FirstName, SecondName: u.SecondName})
	}

	return result, nil
}
