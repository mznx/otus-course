package dialog_message

type DialogMessage struct {
	ID        string `db:"id"`
	DialogID  string `db:"dialog_id"`
	SenderID  string `db:"sender_id"`
	Message   string `db:"message"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
