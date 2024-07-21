package handler

import (
	"net/http"

	"github.com/JacobRWebb/InventoryManagement/internal/service"
	"github.com/a-h/templ"
)

type Handler struct {
	UserHandler      UserHandler
	DashboardHandler DashboardHandler
}

type DashboardHandler interface {
	HandleDashboardGet(http.ResponseWriter, *http.Request)
}

type UserHandler interface {
	HandleUserCreateGet(http.ResponseWriter, *http.Request)
	HandleUserCreatePost(http.ResponseWriter, *http.Request)
	HandleUserLoginGet(http.ResponseWriter, *http.Request)
	HandleUserLoginPost(http.ResponseWriter, *http.Request)
	HandleUserLogoutGet(http.ResponseWriter, *http.Request)
	HandleUserLogoutPost(http.ResponseWriter, *http.Request)
}

func NewHandler(services *service.Service) *Handler {
	h := &Handler{
		UserHandler:      NewUserHandler(services),
		DashboardHandler: NewDashboardHandler(services),
	}
	return h
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}

func HxRedirect(w http.ResponseWriter, r *http.Request, location string) {
	w.Header().Set("HX-Redirect", location)
	w.WriteHeader((http.StatusOK))
}
