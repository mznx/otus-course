package post

import "context"

type Repository interface {
	FindById(ctx context.Context, postId string) (*Post, error)

	Find(ctx context.Context, authorsId []string, offset uint64, limit uint64) ([]*Post, error)

	Update(ctx context.Context, post *Post) error

	Delete(ctx context.Context, postId string) error

	Create(ctx context.Context, post *Post) error
}
