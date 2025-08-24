package user

type GetByIdResponse struct {
	UserID     string `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

type SearchResponse struct {
	UserID     string `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}
