package main

import (
	"log"

	"github.com/JacobRWebb/InventoryManagement/pkg/consul"
	"github.com/JacobRWebb/InventoryManagement/pkg/handlers"
	"github.com/JacobRWebb/InventoryManagement/pkg/middlewares"
	"github.com/JacobRWebb/InventoryManagement/pkg/server"
	"github.com/JacobRWebb/InventoryManagement/pkg/store"

	"github.com/JacobRWebb/InventoryManagement/pkg/config"
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

	store, err := store.NewStore(cfg, consulClient)

	if err != nil {
		log.Fatalf("There was an issue creating stores. `%v`", err)
	}

	handlers := handlers.NewHandler(store)

	middlewares := middlewares.NewMiddleware(store)

	_ = server.NewServer(cfg, handlers, middlewares)
}
