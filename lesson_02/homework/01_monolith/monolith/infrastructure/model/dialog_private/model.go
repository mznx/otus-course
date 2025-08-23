package dialog_private

type DialogPrivate struct {
	UserPairHash string `db:"user_pair_hash"`
	DialogID     string `db:"dialog_id"`
	UserID1      string `db:"user_id_1"`
	UserID2      string `db:"user_id_2"`
	CreatedAt    string `db:"created_at"`
	UpdatedAt    string `db:"updated_at"`
}
