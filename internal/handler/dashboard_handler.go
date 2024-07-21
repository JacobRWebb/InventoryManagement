package handler

import (
	"net/http"

	"github.com/JacobRWebb/InventoryManagement/internal/service"
	session "github.com/JacobRWebb/InventoryManagement/internal/util"
	"github.com/JacobRWebb/InventoryManagement/internal/web/templates/pages"
)

type dashboardHandler struct {
	services *service.Service
}

func NewDashboardHandler(services *service.Service) DashboardHandler {
	return &dashboardHandler{
		services: services,
	}
}

func (h *dashboardHandler) HandleDashboardGet(w http.ResponseWriter, r *http.Request) {
	authRes, err := session.GetAuthResponse(r)

	if authRes == nil || err != nil {
		session.ClearAuthResponse(w, r)
		Render(w, r, pages.Index(nil))
		return
	}

	userId, err := h.services.UserService.ValidateToken(authRes.AccessToken)

	if err != nil {
		session.ClearAuthResponse(w, r)
		Render(w, r, pages.Index(nil))
		return
	}

	user, err := h.services.UserService.GetUser(userId)

	if user == nil || err != nil {
		session.ClearAuthResponse(w, r)
		Render(w, r, pages.Index(nil))
		return
	}

	Render(w, r, pages.Index(user))
}
