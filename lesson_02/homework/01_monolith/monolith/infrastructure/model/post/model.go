package post

type Post struct {
	ID        string `db:"id"`
	AuthorID  string `db:"author_id"`
	Text      string `db:"text"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
