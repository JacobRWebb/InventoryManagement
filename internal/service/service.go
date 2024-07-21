package service

import (
	"github.com/JacobRWebb/InventoryManagement/internal/client"
	"github.com/JacobRWebb/InventoryManagement/internal/models"
)

type Service struct {
	UserService UserService
}

type UserService interface {
	RegisterUser(email string, password string) error
	LoginUser(email string, password string) (*models.AuthResponse, error)
	ValidateToken(token string) (string, error)
	GetUser(userId string) (*models.User, error)
}

func NewService(clients *client.ProtoClients) *Service {
	s := &Service{
		UserService: NewUserService(clients),
	}
	return s
}
