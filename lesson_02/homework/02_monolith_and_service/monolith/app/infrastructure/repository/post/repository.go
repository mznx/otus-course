package post

import (
	"context"
	"database/sql"
	"monolith/domain/post"
	post_mapper "monolith/infrastructure/mapper/post"
	post_model "monolith/infrastructure/model/post"

	"github.com/jmoiron/sqlx"
)

type PostPgRepository struct {
	db *sqlx.DB
}

func NewPostPgRepository(db *sqlx.DB) *PostPgRepository {
	return &PostPgRepository{
		db: db,
	}
}

func (r *PostPgRepository) FindById(ctx context.Context, postId string) (*post.Post, error) {
	var post post_model.Post

	err := r.db.GetContext(ctx, &post, "SELECT * FROM posts WHERE id=$1", postId)

	switch err {
	case nil:
		return post_mapper.ModelToPost(post), nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

func (r *PostPgRepository) Find(ctx context.Context, authorsId []string, offset uint64, limit uint64) ([]*post.Post, error) {
	var posts []post_model.Post

	query, args, err := sqlx.In("SELECT * FROM posts WHERE author_id IN (?) ORDER BY created_at DESC LIMIT ? OFFSET ?", authorsId, limit, offset)
	if err != nil {
		return nil, err
	}

	query = r.db.Rebind(query)

	if err := r.db.SelectContext(ctx, &posts, query, args...); err != nil {
		return nil, err
	}

	return post_mapper.ModelsToPosts(posts), nil
}

func (r *PostPgRepository) Update(ctx context.Context, post *post.Post) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE posts
		SET author_id=$1, text=$2
		WHERE id=$3`, post.AuthorID, post.Text, post.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostPgRepository) Delete(ctx context.Context, postId string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM posts WHERE id=$1", postId)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostPgRepository) Create(ctx context.Context, post *post.Post) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO posts (id, author_id, text)
		VALUES ($1, $2, $3)`, post.ID, post.AuthorID, post.Text)

	if err != nil {
		return err
	}

	return nil
}
