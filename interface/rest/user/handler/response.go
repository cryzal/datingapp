package handler

type DefaultResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type LoginResponse struct {
	Token string
}
