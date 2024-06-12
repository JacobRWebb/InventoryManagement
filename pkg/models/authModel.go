package models

type LoginAccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserByEmailRequest struct {
	Email string `json:"email" validate:"required,email,lte=255"`
}
