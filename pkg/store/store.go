package store

import (
	UserProto "github.com/JacobRWebb/InventoryManagement.Users.Api/pkg/api"
	"github.com/JacobRWebb/InventoryManagement/pkg/config"
	"github.com/JacobRWebb/InventoryManagement/pkg/consul"
	grpcprotoclients "github.com/JacobRWebb/InventoryManagement/pkg/grpc_protoclients"
	"github.com/JacobRWebb/InventoryManagement/pkg/models"
)

type Store struct {
	UserStore UserStore
}

type ProtoClients struct {
	UserServiceClient UserProto.UserServiceClient
}

type UserStore interface {
	RegisterUser(email string, password string) (*models.AuthResponse, error)
	LoginUser(email string, password string) (*models.AuthResponse, error)
}

func NewStore(cfg *config.Config, client *consul.Client, protoClients *grpcprotoclients.ProtoClients) (*Store, error) {
	userStore, err := NewUserStore(cfg, client, protoClients)

	if err != nil {
		panic(err)
	}

	return &Store{
		UserStore: userStore,
	}, nil
}
