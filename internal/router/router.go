package router

import (
	"github.com/JacobRWebb/InventoryManagement/internal/handler"
	"github.com/go-chi/chi/v5"
)

func NewRouter(handlers *handler.Handler) *chi.Mux {
	r := chi.NewRouter()

	AttachMiddlewareRoutes(r)

	AttachStaticRoutes(r)

	r.Get("/login", handlers.UserHandler.HandleUserLoginGet)

	return r
}
