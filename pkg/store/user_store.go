package store

import (
	"fmt"

	"github.com/JacobRWebb/InventoryManagement/pkg/config"
	"github.com/JacobRWebb/InventoryManagement/pkg/consul"
	UserServiceProto "github.com/JacobRWebb/InventoryManagement/pkg/proto/v1/user"
	"google.golang.org/grpc"
)

type userStore struct {
	cfg         *config.Config
	client      *consul.Client
	grpcConn    *grpc.ClientConn
	protoClient UserServiceProto.UserServiceClient
}

func NewUserStore(cfg *config.Config, client *consul.Client) (UserStore, error) {
	us := &userStore{
		cfg:    cfg,
		client: client,
	}

	if err := us.init(); err != nil {
		return nil, err
	}

	return us, nil
}

func (us *userStore) init() error {
	grpcTarget, err := FetchServiceInfo(us.client, us.cfg.UserServiceName)
	if err != nil {
		return fmt.Errorf("fetching service info: %w", err)
	}

	us.grpcConn, err = ConnectGRPC(grpcTarget)
	if err != nil {
		return fmt.Errorf("connecting to gRPC: %w", err)
	}

	us.protoClient = UserServiceProto.NewUserServiceClient(us.grpcConn)

	return nil
}

// func (us *UserStore) CreateUser(email string, password string) (string, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
// 	defer cancel()

// 	newUser := &UserServiceProto.CreateUserRequest{
// 		Email:    email,
// 		Password: password,
// 	}

// 	response, err := us.protoClient.CreateUser(ctx, newUser)

// 	if err != nil {
// 		return "", fmt.Errorf("%v - %v", USER_CREATION_ERROR, err)
// 	}

// 	switch result := response.Result.(type) {
// 	case *UserServiceProto.CreateUserResponse_UserId:
// 		return result.UserId, nil
// 	case *UserServiceProto.CreateUserResponse_Error:
// 		return "", USER_CREATION_ERROR
// 	default:
// 		return "", USER_CREATION_ERROR
// 	}
// }

// func (us *UserStore) CheckIfEmailExists(email string) bool {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
// 	defer cancel()

// 	response, err := us.protoClient.CheckEmailExist(ctx, &UserServiceProto.CheckEmailExistRequest{Email: email})

// 	if err != nil {
// 		fmt.Printf("There was an error %v", err)
// 	}

// 	return response.Exists
// }
