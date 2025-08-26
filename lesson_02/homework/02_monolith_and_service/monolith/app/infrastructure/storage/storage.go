package storage

import (
	"monolith/domain/post"
	"monolith/domain/user"
)

type Repository struct {
	User user.Repository
	Post post.Repository
}
