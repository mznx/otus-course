package friend

type Friend struct {
	UserID    string `db:"user_id"`
	FriendID  string `db:"friend_id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
