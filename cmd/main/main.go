package main

import (
	"encoding/gob"
	"log"

	"github.com/JacobRWebb/InventoryManagement/pkg/consul"
	grpcprotoclients "github.com/JacobRWebb/InventoryManagement/pkg/grpc_protoclients"
	"github.com/JacobRWebb/InventoryManagement/pkg/handlers"
	"github.com/JacobRWebb/InventoryManagement/pkg/middlewares"
	"github.com/JacobRWebb/InventoryManagement/pkg/models"
	"github.com/JacobRWebb/InventoryManagement/pkg/server"
	"github.com/JacobRWebb/InventoryManagement/pkg/store"
	"github.com/gorilla/sessions"

	"github.com/JacobRWebb/InventoryManagement/pkg/config"
)

func main() {

	gob.Register(&models.AuthResponse{})

	cookieSession := sessions.NewCookieStore([]byte("Secret-Key"))

	cookieSession.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   true,
	}

	middlewares.SessionStore = cookieSession
	if middlewares.SessionStore == nil {
		log.Fatal("SessionStore is nil after assignment")
	}

	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	consulClient, err := consul.NewClient(cfg.ConsulAddr)

	if err != nil {
		log.Fatalf("Creating Consul Client error: %v", err)
	}

	protoClients := grpcprotoclients.NewProtoClients(cfg, consulClient)

	store, err := store.NewStore(cfg, consulClient, protoClients)

	if err != nil {
		log.Fatalf("There was an issue creating stores. `%v`", err)
	}

	handlers := handlers.NewHandler(store)

	middlewares := middlewares.NewMiddleware(store)

	_ = server.NewServer(cfg, handlers, middlewares)
}
