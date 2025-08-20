package user_auth

type UserAuth struct {
	UserID    string `db:"user_id"`
	PassHash  string `db:"pass_hash"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
