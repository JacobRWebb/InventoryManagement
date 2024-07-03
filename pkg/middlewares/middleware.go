package middlewares

import (
	"net/http"

	"github.com/JacobRWebb/InventoryManagement/pkg/store"
)

type Middleware struct {
	UserMiddleware UserMiddleware
}

type UserMiddleware interface {
	AuthMiddleware(next http.HandlerFunc) http.HandlerFunc
}

func NewMiddleware(store *store.Store) *Middleware {
	m := &Middleware{
		UserMiddleware: NewUserMiddleware(store),
	}

	return m
}
