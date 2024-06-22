package consul

import (
	"errors"
	"fmt"

	"github.com/hashicorp/consul/api"
)

type Client struct {
	client *api.Client
}

func NewClient(addr string) (*Client, error) {
	config := api.DefaultConfig()
	config.Address = addr

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

// func (c *Client) updateHealthChecker() {
// 	ticker := time.NewTicker(time.Second * 5)
// 	for {
// 		c.client.Agent().UpdateTTL("InventoryManagement-CheckAlive", "online", api.HealthPassing)
// 		<-ticker.C
// 	}
// }

// func (c *Client) Register(cfg *config.Config) (err error) {
// 	log.Printf("%s:%d", cfg.GRPCAddr, cfg.GRPCPort)
// 	// GRPC Service Consul
// 	grpcReg := &api.AgentServiceRegistration{
// 		ID:      "InventoryManagement",
// 		Name:    cfg.ServiceName,
// 		Address: cfg.ConsulAddr,
// 		Port:    cfg.GRPCPort,
// 		Check: &api.AgentServiceCheck{
// 			TTL:                            fmt.Sprintf("%ds", 8),
// 			DeregisterCriticalServiceAfter: cfg.ConsulDeregisterTime,
// 			CheckID:                        "InventoryManagement-CheckAlive",
// 		},
// 	}

// 	err = c.client.Agent().ServiceRegister(grpcReg)

// 	go c.updateHealthChecker()

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
