package store

import (
	"context"
	"fmt"
	"time"

	UserProto "github.com/JacobRWebb/InventoryManagement.Users.Api/pkg/api"
	"github.com/JacobRWebb/InventoryManagement/pkg/config"
	"github.com/JacobRWebb/InventoryManagement/pkg/consul"
	grpcprotoclients "github.com/JacobRWebb/InventoryManagement/pkg/grpc_protoclients"
	"github.com/JacobRWebb/InventoryManagement/pkg/models"
)

type userStore struct {
	cfg          *config.Config
	client       *consul.Client
	protoClients *grpcprotoclients.ProtoClients
}

func NewUserStore(cfg *config.Config, client *consul.Client, protoClients *grpcprotoclients.ProtoClients) (UserStore, error) {
	us := &userStore{
		cfg:          cfg,
		client:       client,
		protoClients: protoClients,
	}

	return us, nil
}

func (us *userStore) RegisterUser(email string, password string) (*models.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	registerUserRequest := &UserProto.RegisterUserRequest{
		Email:    email,
		Password: password,
	}

	response, err := us.protoClients.UserServiceClient.RegisterUser(ctx, registerUserRequest)

	if err != nil {
		return nil, err
	}

	ar := &models.AuthResponse{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
		TokenType:    response.TokenType,
	}

	return ar, nil
}

func (us *userStore) LoginUser(email string, password string) (*models.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	loginUserRequest := &UserProto.LoginUserRequest{
		Email:    email,
		Password: password,
	}

	response, err := us.protoClients.UserServiceClient.LoginUser(ctx, loginUserRequest)

	if err != nil {
		return nil, err
	}

	ar := &models.AuthResponse{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
		TokenType:    response.TokenType,
	}

	fmt.Printf("%v", ar)

	return ar, nil
}
