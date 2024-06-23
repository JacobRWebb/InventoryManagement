package server

import (
	"github.com/JacobRWebb/InventoryManagement/pkg/config"
	"github.com/JacobRWebb/InventoryManagement/pkg/store"
)

type Server struct {
	cfg *config.Config
	// router *chi.Mux
	store *store.Store
}

func NewServer(cfg *config.Config, store *store.Store) *Server {
	s := &Server{
		cfg:   cfg,
		store: store,
	}

	s.routes()

	return s
}

func (s *Server) routes() {

}
