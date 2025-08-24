package post

import "github.com/google/uuid"

type Post struct {
	ID       string
	Text     string
	AuthorID string
}

func NewPost(authorId string, text string) *Post {
	return &Post{
		ID:       uuid.New().String(),
		AuthorID: authorId,
		Text:     text,
	}
}
