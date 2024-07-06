package middlewares

import (
	"net/http"

	"github.com/JacobRWebb/InventoryManagement/pkg/store"
	"github.com/gorilla/sessions"
)

var SessionStore *sessions.CookieStore

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
