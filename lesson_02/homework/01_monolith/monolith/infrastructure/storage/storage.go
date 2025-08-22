package storage

import (
	"monolith/domain/user"
)

type Repository struct {
	User user.Repository
}
