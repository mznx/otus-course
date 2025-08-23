package dialog_member

type DialogMember struct {
	UserID    string `db:"user_id"`
	DialogID  string `db:"dialog_id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
