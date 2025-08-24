package post

import (
	"monolith/domain/post"
	post_model "monolith/infrastructure/model/post"
)

func ModelToPost(p post_model.Post) *post.Post {
	return &post.Post{
		ID:       p.ID,
		AuthorID: p.AuthorID,
		Text:     p.Text,
	}
}

func ModelsToPosts(posts []post_model.Post) []*post.Post {
	result := []*post.Post{}

	for _, p := range posts {
		result = append(result, ModelToPost(p))
	}

	return result
}
