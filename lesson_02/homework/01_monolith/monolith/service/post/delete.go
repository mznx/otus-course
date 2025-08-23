package post

import (
	"context"
	"errors"
	"monolith/domain/post"
)

type DeleteData struct {
	PostID string
}

type PostDeleteService struct {
	postRepository post.Repository
}

func NewPostDeleteService(postRepository post.Repository) *PostDeleteService {
	return &PostDeleteService{postRepository: postRepository}
}

func (s *PostDeleteService) Handle(ctx context.Context, data *DeleteData) error {
	post, err := s.postRepository.FindById(ctx, data.PostID)

	if err != nil {
		return err
	}

	if post == nil {
		return errors.New("post not found")
	}

	if err = s.postRepository.Delete(ctx, post.ID); err != nil {
		return err
	}

	return nil
}
