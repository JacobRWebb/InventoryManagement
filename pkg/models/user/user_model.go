package user_models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID `gorm:"Id" json:"id"`
	Email        string    `gorm:"unique" json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	ModifiedAt   time.Time `json:"modified_at"`
}

type CreateUserRequest struct {
	Email           string `json:"email" validate:"required,email,lte=255"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}
