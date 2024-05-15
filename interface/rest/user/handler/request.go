package handler

type CreatePayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
