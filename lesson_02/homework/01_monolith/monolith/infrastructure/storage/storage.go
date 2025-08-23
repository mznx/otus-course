package storage

import (
	"monolith/domain/dialog"
	"monolith/domain/user"
)

type Repository struct {
	User   user.Repository
	Dialog dialog.Repository
}
