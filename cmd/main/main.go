package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/JacobRWebb/InventoryManagement/pkg/consul"
	UserServiceProto "github.com/JacobRWebb/InventoryManagement/pkg/proto/v1/user"

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

	// err = consulClient.Register(cfg)

	// if err != nil {
	// 	log.Fatalf("Consul Client error: %v", err)
	// }

	addr, port, err := consulClient.FindService("User_Service")

	if err != nil {
		log.Fatalf("Consul service discovery error: %v", err)
	}

	target := fmt.Sprintf("%s:%d", addr, port)

	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("%v", err)
	}

	client := UserServiceProto.NewServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	req := &UserServiceProto.CreateUserRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	resp, err := client.CreateUser(ctx, req)
	if err != nil {
		log.Fatalf("CreateUser request failed: %v", err)
	}

	log.Printf("CreateUser response: %v", resp)
}

// conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", cfg.GRPCPort), grpc.WithInsecure(), grpc.WithBlock())
// if err != nil {
// 	log.Fatalf("Failed to connect to gRPC server: %v", err)
// }
// defer conn.Close()

// client := pb.NewServiceClient(conn)

// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
