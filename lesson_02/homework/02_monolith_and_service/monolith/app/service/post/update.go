package post

import (
	"context"
	"errors"
	"monolith/domain/post"
)

type UpdateData struct {
	PostID string
	Text   string
}

type PostUpdateService struct {
	postRepository post.Repository
}

func NewPostUpdateService(postRepository post.Repository) *PostUpdateService {
	return &PostUpdateService{postRepository: postRepository}
}

func (s *PostUpdateService) Handle(ctx context.Context, data *UpdateData) error {
	post, err := s.postRepository.FindById(ctx, data.PostID)

	if err != nil {
		return err
	}

	if post == nil {
		return errors.New("post not found")
	}

	post.Text = data.Text

	if err = s.postRepository.Update(ctx, post); err != nil {
		return err
	}

	return nil
}
