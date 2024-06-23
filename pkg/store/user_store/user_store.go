package userstore

import (
	"context"
	"fmt"
	"time"

	"github.com/JacobRWebb/InventoryManagement/pkg/config"
	"github.com/JacobRWebb/InventoryManagement/pkg/consul"
	UserServiceProto "github.com/JacobRWebb/InventoryManagement/pkg/proto/v1/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	USER_CREATION_ERROR = fmt.Errorf("there was a problem creating the user")
)

type UserStore struct {
	cfg         *config.Config
	client      *consul.Client
	grpcTarget  string
	protoClient UserServiceProto.UserServiceClient
}

func NewUserStore(cfg *config.Config, client *consul.Client) (us *UserStore, err error) {
	us = &UserStore{
		cfg:    cfg,
		client: client,
	}

	err = us.fetchServiceInfo()

	if err != nil {
		return nil, err
	}

	err = us.connectGrpc()

	if err != nil {
		return nil, err
	}

	return us, nil
}

func (s *UserStore) fetchServiceInfo() error {
	addr, port, err := s.client.FindService(s.cfg.UserServiceName)

	if err != nil {
		return fmt.Errorf("there was an issue finding consul %s", s.cfg.UserServiceName)
	}

	s.grpcTarget = fmt.Sprintf("%s:%d", addr, port)

	return nil
}

func (s *UserStore) connectGrpc() error {
	conn, err := grpc.NewClient(s.grpcTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return err
	}

	s.protoClient = UserServiceProto.NewUserServiceClient(conn)

	return nil
}

func (us *UserStore) CreateUser(email string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	newUser := &UserServiceProto.CreateUserRequest{
		Email:    email,
		Password: password,
	}

	response, err := us.protoClient.CreateUser(ctx, newUser)

	if err != nil {
		return "", fmt.Errorf("%v - %v", USER_CREATION_ERROR, err)
	}

	switch result := response.Result.(type) {
	case *UserServiceProto.CreateUserResponse_UserId:
		return result.UserId, nil
	case *UserServiceProto.CreateUserResponse_Error:
		return "", USER_CREATION_ERROR
	default:
		return "", USER_CREATION_ERROR
	}
}
