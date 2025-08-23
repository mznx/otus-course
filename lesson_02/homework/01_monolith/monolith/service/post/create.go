package post

import (
	"context"
	"monolith/domain/post"
)

type PostCreateData struct {
	UserID string
	Text   string
}

type PostCreateResult struct {
	PostID string
}

type PostCreateService struct {
	postRepository post.Repository
}

func NewPostCreateService(postRepository post.Repository) *PostCreateService {
	return &PostCreateService{postRepository: postRepository}
}

func (s *PostCreateService) Handle(ctx context.Context, data *PostCreateData) (*PostCreateResult, error) {
	post := post.NewPost(data.UserID, data.Text)

	if err := s.postRepository.Create(ctx, post); err != nil {
		return nil, err
	}

	return &PostCreateResult{PostID: post.ID}, nil
}
