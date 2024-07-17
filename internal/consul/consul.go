package consul

import (
	"errors"
	"fmt"

	"github.com/JacobRWebb/InventoryManagement/internal/config"
	"github.com/hashicorp/consul/api"
)

type Client struct {
	client *api.Client
}

func NewClient(cfg *config.Config) (*Client, error) {
	config := api.DefaultConfig()
	config.Address = cfg.ConsulConfig.Address

	client, err := api.NewClient(config)

	if err != nil {
		return nil, fmt.Errorf("failed to create Consule client: %v", err)
	}

	return &Client{client: client}, nil
}

func (s *Client) FindService(serviceName string) (serviceAddr string, servicePort int, err error) {
	services, _, err := s.client.Catalog().Service(serviceName, "", nil)

	if err != nil {
		return "", -1, err
	}

	if len(services) == 0 {
		return "", -1, errors.New("no services found")
	}

	service := services[0]

	address := serviceAddr
	port := service.ServicePort

	return address, port, nil
}
