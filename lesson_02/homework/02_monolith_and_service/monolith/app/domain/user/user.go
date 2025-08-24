package user

import "github.com/google/uuid"

type User struct {
	ID         string
	FirstName  string
	SecondName string
}

func NewUser(firstName string, secondName string) *User {
	return &User{
		ID:         uuid.New().String(),
		FirstName:  firstName,
		SecondName: secondName,
	}
}
