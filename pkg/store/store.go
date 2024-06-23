package store

import (
	"github.com/JacobRWebb/InventoryManagement/pkg/config"
	"github.com/JacobRWebb/InventoryManagement/pkg/consul"
	userstore "github.com/JacobRWebb/InventoryManagement/pkg/store/user_store"
)

type Store struct {
	UserStore *userstore.UserStore
}

func NewStore(cfg *config.Config, client *consul.Client) (*Store, error) {

	us, err := userstore.NewUserStore(cfg, client)

	if err != nil {
		return nil, err
	}

	return &Store{
		UserStore: us,
	}, nil
}
