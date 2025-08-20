package user

import "github.com/google/uuid"

type User struct {
	ID         string
	FirstName  string
	SecondName string
}

func NewUser(firstName string, secondName string) *User {
	id := uuid.New()

	return &User{
		ID:         id.String(),
		FirstName:  firstName,
		SecondName: secondName,
	}
}
