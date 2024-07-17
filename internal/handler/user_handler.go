package handler

import (
	"net/http"

	"github.com/JacobRWebb/InventoryManagement/internal/models"
	"github.com/JacobRWebb/InventoryManagement/internal/service"
	"github.com/JacobRWebb/InventoryManagement/internal/web/templates/pages"
)

type userHandler struct {
	services *service.Service
}

func NewUserHandler(services *service.Service) UserHandler {
	h := &userHandler{
		services: services,
	}

	return h
}

func (h *userHandler) HandleUserLoginGet(w http.ResponseWriter, r *http.Request) {
	values, _ := models.ParseLoginAccountFormValuesAndValidate(r)

	Render(w, r, pages.Login(values))
}
