package service

import "github.com/JacobRWebb/InventoryManagement/internal/client"

type Service struct {
	UserService UserService
}

type UserService interface {
}

func NewService(clients *client.ProtoClients) *Service {
	s := &Service{
		UserService: NewUserService(),
	}
	return s
}
