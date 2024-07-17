package handler

import (
	"net/http"

	"github.com/JacobRWebb/InventoryManagement/internal/service"
	"github.com/a-h/templ"
)

type Handler struct {
	UserHandler UserHandler
}

type UserHandler interface {
	HandleUserLoginGet(http.ResponseWriter, *http.Request)
}

func NewHandler(services *service.Service) *Handler {
	h := &Handler{
		UserHandler: NewUserHandler(services),
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
