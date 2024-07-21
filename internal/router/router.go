package router

import (
	"github.com/JacobRWebb/InventoryManagement/internal/handler"
	"github.com/JacobRWebb/InventoryManagement/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter(handlers *handler.Handler) *chi.Mux {
	r := chi.NewRouter()

	m := middleware.NewMiddleware(r, handlers)
	m.RunGeneric()
	m.ServeStatic()

	r.Group(func(r chi.Router) {
		r.Use(m.AuthMiddleware.ProtectedRoute)

		r.Get("/", handlers.DashboardHandler.HandleDashboardGet)
		r.Get("/logout", handlers.UserHandler.HandleUserLogoutGet)
		r.Post("/logout", handlers.UserHandler.HandleUserLogoutPost)
	})

	r.Group(func(r chi.Router) {
		r.Use(m.AuthMiddleware.UnprotectedRoute)
		r.Get("/login", handlers.UserHandler.HandleUserLoginGet)
		r.Post("/login", handlers.UserHandler.HandleUserLoginPost)

		r.Get("/register", handlers.UserHandler.HandleUserCreateGet)
		r.Post("/register", handlers.UserHandler.HandleUserCreatePost)
	})

	return r
}
