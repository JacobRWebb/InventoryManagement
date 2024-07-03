package handlers

import (
	"net/http"

	"github.com/JacobRWebb/InventoryManagement/pkg/store"
	"github.com/a-h/templ"
)

type Handles struct {
	UserHandler      UserHandler
	BasicHandler     BasicHandler
	DashboardHandler DashboardHandler
}

type BasicHandler interface {
	HandleStatic(http.Handler) http.HandlerFunc
	HandleFaviconGet(http.ResponseWriter, *http.Request)
}

type UserHandler interface {
	HandleUserCreateGet(http.ResponseWriter, *http.Request)
	HandleUserCreatePost(http.ResponseWriter, *http.Request)
	HandleUserLoginGet(http.ResponseWriter, *http.Request)
	HandleUserLoginPost(http.ResponseWriter, *http.Request)
	HandleUserLogoutPost(http.ResponseWriter, *http.Request)
}

type DashboardHandler interface {
	HandleDashboardGet(w http.ResponseWriter, r *http.Request)
}

func NewHandler(store *store.Store) *Handles {
	return &Handles{
		UserHandler:      NewUserHandler(store),
		BasicHandler:     NewBasicHandler(),
		DashboardHandler: NewDashboardHandler(store),
	}
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}

func HxRedirect(w http.ResponseWriter, r *http.Request, location string) {
	w.Header().Set("HX-Redirect", location)
	w.WriteHeader((http.StatusOK))
}
