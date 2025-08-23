package post

import "github.com/google/uuid"

type Post struct {
	ID       string
	Text     string
	AuthorID string
}

func NewPost(authorId string, text string) *Post {
	id := uuid.New()

	return &Post{
		ID:       id.String(),
		AuthorID: authorId,
		Text:     text,
	}
}
