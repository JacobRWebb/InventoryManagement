package models

type RegisterAccountRequest struct {
	Email           string `json:"email" validate:"required,email,lte=255"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type LoginAccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserByEmailRequest struct {
	Email string `json:"email" validate:"required,email,lte=255"`
}
