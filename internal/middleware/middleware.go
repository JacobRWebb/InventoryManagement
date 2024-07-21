package middleware

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/JacobRWebb/InventoryManagement/internal/handler"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

type Middleware struct {
	r              *chi.Mux
	AuthMiddleware AuthMiddleware
}

type AuthMiddleware interface {
	ProtectedRoute(next http.Handler) http.Handler
	UnprotectedRoute(next http.Handler) http.Handler
}

func NewMiddleware(r *chi.Mux, handlers *handler.Handler) *Middleware {
	return &Middleware{
		r:              r,
		AuthMiddleware: NewAuthMiddleware(r, handlers),
	}
}

func (m *Middleware) RunGeneric() {
	m.r.Use(chiMiddleware.RequestID)
	m.r.Use(chiMiddleware.RealIP)
	m.r.Use(chiMiddleware.Logger)
	m.r.Use(chiMiddleware.Recoverer)
	m.r.Use(chiMiddleware.Timeout(60 * time.Second))
}

func (m *Middleware) ServeStatic() {
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "/internal/web/static"))

	m.r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(filesDir)))
	m.r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(workDir, "/internal/web/static/assets/favicon.ico"))
	})
}
