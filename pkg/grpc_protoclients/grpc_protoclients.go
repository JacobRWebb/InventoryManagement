package grpcprotoclients

import (
	"fmt"

	UserProto "github.com/JacobRWebb/InventoryManagement.Users.Api/pkg/api"
	"github.com/JacobRWebb/InventoryManagement/pkg/config"
	"github.com/JacobRWebb/InventoryManagement/pkg/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProtoClients struct {
	UserServiceClient UserProto.UserServiceClient
}

func NewProtoClients(cfg *config.Config, consulClient *consul.Client) *ProtoClients {

	userServiceClient, err := NewUserServiceClient(cfg, consulClient)

	if err != nil {
		panic(err)
	}

	protoClients := &ProtoClients{
		UserServiceClient: userServiceClient,
	}

	return protoClients
}

func FetchServiceInfo(consulClient *consul.Client, serviceName string) (string, error) {
	addr, port, err := consulClient.FindService(serviceName)
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
