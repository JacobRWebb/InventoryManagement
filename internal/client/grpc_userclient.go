package client

import (
	"fmt"

	UserProto "github.com/JacobRWebb/InventoryManagement.Users.Api/pkg/api"
	"github.com/JacobRWebb/InventoryManagement/internal/config"
	"github.com/JacobRWebb/InventoryManagement/internal/consul"
)

func NewUserServiceClient(cfg *config.Config, client *consul.Client) (UserProto.UserServiceClient, error) {
	grpcTarget, err := FetchServiceInfo(client, cfg.ServiceConfig.UserServiceName)
	if err != nil {
		return nil, fmt.Errorf("fetching service info: %w", err)
	}

	grpcConn, err := ConnectGRPC(grpcTarget)

	if err != nil {
		return nil, fmt.Errorf("connecting to gRPC: %w", err)
	}

	return UserProto.NewUserServiceClient(grpcConn), nil
}
