package helper

import "github.com/google/uuid"

func GenerateAuthToken(userId string) string {
	token := uuid.New() // generate token

	return token.String()
}
