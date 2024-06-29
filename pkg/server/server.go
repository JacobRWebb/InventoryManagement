package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/JacobRWebb/InventoryManagement/pkg/config"
	"github.com/JacobRWebb/InventoryManagement/pkg/store"
	"github.com/JacobRWebb/InventoryManagement/pkg/web/templates/pages"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/net/http2"
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

	fmt.Printf("Server Running https://localhost:%s\n", "3333")

	cert, err := tls.LoadX509KeyPair("certs/pub.pem", "certs/key.pem")
	if err != nil {
		log.Fatalf("Error loading certificate: %v", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	srv := &http.Server{
		Addr:      ":3333",
		Handler:   r,
		TLSConfig: tlsConfig,
	}

	http2.ConfigureServer(srv, &http2.Server{})
	log.Fatal(srv.ListenAndServeTLS("", ""))

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
		pages.Index().Render(context.Background(), w)
	})
}
