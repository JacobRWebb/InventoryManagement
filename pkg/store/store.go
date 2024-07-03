package store

import (
	"fmt"

	"github.com/JacobRWebb/InventoryManagement/pkg/config"
	"github.com/JacobRWebb/InventoryManagement/pkg/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Store struct {
	UserStore UserStore
}

type UserStore interface {
}

func NewStore(cfg *config.Config, client *consul.Client) (*Store, error) {
	userStore, err := NewUserStore(cfg, client)

	if err != nil {
		return nil, fmt.Errorf("creating user store: %v", err)
	}

	return &Store{
		UserStore: userStore,
	}, nil
}

func FetchServiceInfo(client *consul.Client, serviceName string) (string, error) {
	addr, port, err := client.FindService(serviceName)
	if err != nil {
		return "", fmt.Errorf("finding service %s: %w", serviceName, err)
	}
	return fmt.Sprintf("%s:%d", addr, port), nil
}

func ConnectGRPC(target string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("connecting to gRPC server: %w", err)
	}
	return conn, nil
}
