package user

type User struct {
	ID string
}

func NewUser(userId string) *User {
	return &User{
		ID: userId,
	}
}
