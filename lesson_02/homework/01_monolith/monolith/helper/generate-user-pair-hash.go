package helper

func GenerateUserPairHash(userId1 string, userId2 string) string {
	if userId1 < userId2 {
		return userId1 + "." + userId2
	} else {
		return userId2 + "." + userId1
	}
}
