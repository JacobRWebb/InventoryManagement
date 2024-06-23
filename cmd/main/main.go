package main

import (
	"fmt"
	"log"

	"github.com/JacobRWebb/InventoryManagement/pkg/consul"

	"github.com/JacobRWebb/InventoryManagement/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	consulClient, err := consul.NewClient(cfg.ConsulAddr)

	if err != nil {
		log.Fatalf("Creating Consul Client error: %v", err)
	}

	addr, port, err := consulClient.FindService("User_Service")

	if err != nil {
		log.Fatalf("Consul service discovery error: %v", err)
	}

	target := fmt.Sprintf("%s:%d", addr, port)

	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("%v", err)
	}

	// client := UserServiceProto.NewServiceClient(conn)
	conn.Close()
}
