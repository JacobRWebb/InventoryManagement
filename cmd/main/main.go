package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/JacobRWebb/InventoryManagement/internal/config"
	"github.com/JacobRWebb/InventoryManagement/internal/server"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}

func run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)

	// consulClient, err := consul.NewClient(cfg)
	// if err != nil {
	// 	return fmt.Errorf("error while starting consul client: %v", err)
	// }

	srv, err := server.NewServer(cfg)

	if err != nil {
		return fmt.Errorf("error while creating server: %v", err)
	}

	go func() {
		logger.Println("Starting Server")
		server.Run(srv)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Println("Server exiting")

	return nil
}
