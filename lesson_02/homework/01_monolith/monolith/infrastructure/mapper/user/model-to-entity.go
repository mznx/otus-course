package user_mapper

import (
	"monolith/domain/user"
	user_model "monolith/infrastructure/model/user"
)

func ModelToUser(u user_model.User) *user.User {
	return &user.User{
		ID:         u.ID,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
	}
}
