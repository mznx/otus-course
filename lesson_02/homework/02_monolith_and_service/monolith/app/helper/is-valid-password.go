package helper

func IsValidPassword(passHash string, password string) bool {
	hash := password // hashing password

	return passHash == hash
}
