package auth

type LoginRequest struct {
	UserID   string `json:"id"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Password   string `json:"password"`
}

type RegisterResponse struct {
	UserID string `json:"user_id"`
}
