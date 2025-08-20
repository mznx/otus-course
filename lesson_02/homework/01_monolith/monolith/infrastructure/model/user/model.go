package user

type User struct {
	ID         string `db:"id"`
	FirstName  string `db:"first_name"`
	SecondName string `db:"second_name"`
	CreatedAt  string `db:"created_at"`
	UpdatedAt  string `db:"updated_at"`
}
