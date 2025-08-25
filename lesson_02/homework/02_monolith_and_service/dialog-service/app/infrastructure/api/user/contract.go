package user

type FindByIdResponse struct {
	UserID     string `json:"user_id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}
