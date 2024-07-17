package router

import (
	"net/http"

	"github.com/JacobRWebb/InventoryManagement/internal/handler"
	"github.com/go-chi/chi/v5"
)

func NewRouter(handlers *handler.Handler) *chi.Mux {
	r := chi.NewRouter()

	AttachMiddlewareRoutes(r)

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	AttachStaticRoutes(r)

	r.Get("/login", handlers.UserHandler.HandleUserLoginGet)

	return r
}

// 	s.router.Get("/", s.middlewares.UserMiddleware.AuthMiddleware(s.handlers.DashboardHandler.HandleDashboardGet))
// 	s.router.Get("/login", s.handlers.UserHandler.HandleUserLoginGet)
// 	s.router.Post("/login", s.handlers.UserHandler.HandleUserLoginPost)
// 	s.router.Get("/register", s.handlers.UserHandler.HandleUserCreateGet)
// 	s.router.Post("/register", s.handlers.UserHandler.HandleUserCreatePost)
// 	s.router.Post("/logout", s.handlers.UserHandler.HandleUserLogoutPost)
// }
