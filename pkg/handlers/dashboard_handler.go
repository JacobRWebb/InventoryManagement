package handlers

import (
	"net/http"

	"github.com/JacobRWebb/InventoryManagement/pkg/models"
	"github.com/JacobRWebb/InventoryManagement/pkg/store"
	"github.com/JacobRWebb/InventoryManagement/pkg/web/templates/pages"
)

type dashboardHandler struct {
	store *store.Store
}

func NewDashboardHandler(store *store.Store) DashboardHandler {
	dh := &dashboardHandler{
		store: store,
	}

	return dh
}

func (dh *dashboardHandler) HandleDashboardGet(w http.ResponseWriter, r *http.Request) {
	user := &models.User{
		Username: "Jacob Webb",
		Profile: &models.UserProfile{
			ProfilePic: "https://ui-avatars.com/api/?name=Jacob+Webb",
			Email:      "JacobRWebbkc@gmail.com",
		},
	}

	Render(w, r, pages.Index(*user))
}
