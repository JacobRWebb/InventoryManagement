package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/JacobRWebb/InventoryManagement/pkg/config"
	"github.com/JacobRWebb/InventoryManagement/pkg/store"
	"github.com/JacobRWebb/InventoryManagement/pkg/web/templates"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	cfg   *config.Config
	r     *chi.Mux
	store *store.Store
}

func NewServer(cfg *config.Config, store *store.Store) *Server {
	r := chi.NewRouter()

	s := &Server{
		cfg:   cfg,
		r:     r,
		store: store,
	}

	s.applyMiddleware()
	s.routes()

	fmt.Printf("Server Running http://localhost:%s\n", "3333")

	http.ListenAndServe(fmt.Sprintf(":%s", "3333"), r)

	return s
}

func (s *Server) applyMiddleware() {
	s.r.Use(middleware.RequestID)
	s.r.Use(middleware.RealIP)
	s.r.Use(middleware.Logger)
	s.r.Use(middleware.Recoverer)
	s.r.Use(middleware.Timeout(60 * time.Second))
}

func (s *Server) routes() {
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "/pkg/web/static"))

	s.r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(filesDir)))

	s.r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(workDir, "/pkg/web/static/assets/favicon.ico"))
	})

	s.r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templates.IndexPage().Render(context.Background(), w)
	})
}
