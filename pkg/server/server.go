package server

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/JacobRWebb/InventoryManagement/pkg/config"
	"github.com/JacobRWebb/InventoryManagement/pkg/handlers"
	"github.com/JacobRWebb/InventoryManagement/pkg/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/net/http2"
)

type Server struct {
	cfg         *config.Config
	router      *chi.Mux
	handlers    *handlers.Handles
	middlewares *middlewares.Middleware
}

func NewServer(cfg *config.Config, h *handlers.Handles, middlewares *middlewares.Middleware) *Server {
	r := chi.NewRouter()

	s := &Server{
		cfg:         cfg,
		router:      r,
		handlers:    h,
		middlewares: middlewares,
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
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Timeout(60 * time.Second))
}

func (s *Server) routes() {
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "/pkg/web/static"))

	s.router.Handle("/static/*", http.StripPrefix("/static", http.FileServer(filesDir)))

	s.router.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(workDir, "/pkg/web/static/assets/favicon.ico"))
	})

	s.router.Get("/", s.middlewares.UserMiddleware.AuthMiddleware(s.handlers.DashboardHandler.HandleDashboardGet))
	s.router.Get("/login", s.handlers.UserHandler.HandleUserLoginGet)
	s.router.Post("/login", s.handlers.UserHandler.HandleUserLoginPost)
	s.router.Get("/register", s.handlers.UserHandler.HandleUserCreateGet)
	s.router.Post("/register", s.handlers.UserHandler.HandleUserCreatePost)
	s.router.Post("/logout", s.handlers.UserHandler.HandleUserLogoutPost)
}
