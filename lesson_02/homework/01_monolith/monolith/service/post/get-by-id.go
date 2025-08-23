package post

import (
	"context"
	"errors"
	"monolith/domain/post"
)

type GetByIdData struct {
	PostID string
}

type GetByIdResult struct {
	Post *post.Post
}

type PostGetByIdService struct {
	postRepository post.Repository
}

func NewPostGetByIdService(postRepository post.Repository) *PostGetByIdService {
	return &PostGetByIdService{postRepository: postRepository}
}

func (s *PostGetByIdService) Handle(ctx context.Context, data *GetByIdData) (*GetByIdResult, error) {
	post, err := s.postRepository.FindById(ctx, data.PostID)

	if err != nil {
		return nil, err
	}

	if post == nil {
		return nil, errors.New("post not found")
	}

	return &GetByIdResult{Post: post}, nil
}
