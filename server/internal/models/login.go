package models

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	UserGetAllResponse `json:"user"`
	Token              string `json:"token"`
}
