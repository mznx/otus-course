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

func ModelsToUsers(users []user_model.User) []*user.User {
	result := []*user.User{}

	for _, u := range users {
		result = append(result, ModelToUser(u))
	}

	return result
}
