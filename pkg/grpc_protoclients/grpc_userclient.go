package grpcprotoclients

import (
	"fmt"

	UserProto "github.com/JacobRWebb/InventoryManagement.Users.Api/pkg/api"
	"github.com/JacobRWebb/InventoryManagement/pkg/config"
	"github.com/JacobRWebb/InventoryManagement/pkg/consul"
)

func NewUserServiceClient(cfg *config.Config, client *consul.Client) (UserProto.UserServiceClient, error) {
	grpcTarget, err := FetchServiceInfo(client, cfg.UserServiceName)
	if err != nil {
		return nil, fmt.Errorf("fetching service info: %w", err)
	}

	grpcConn, err := ConnectGRPC(grpcTarget)

	if err != nil {
		return nil, fmt.Errorf("connecting to gRPC: %w", err)
	}

	return UserProto.NewUserServiceClient(grpcConn), nil
}
