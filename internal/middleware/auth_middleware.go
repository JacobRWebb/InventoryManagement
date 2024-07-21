package middleware

import (
	"net/http"

	"github.com/JacobRWebb/InventoryManagement/internal/handler"
	session "github.com/JacobRWebb/InventoryManagement/internal/util"
	"github.com/go-chi/chi/v5"
)

type authMiddleware struct {
	r        *chi.Mux
	handlers *handler.Handler
}

func NewAuthMiddleware(r *chi.Mux, handlers *handler.Handler) AuthMiddleware {
	return &authMiddleware{
		r:        r,
		handlers: handlers,
	}
}

func (m *authMiddleware) ProtectedRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authRes, err := session.GetAuthResponse(r)

		if authRes == nil && err != nil {
			session.ClearAuthResponse(w, r)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *authMiddleware) UnprotectedRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authRes, err := session.GetAuthResponse(r)

		if authRes != nil && err == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
