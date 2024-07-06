package handlers

import (
	"net/http"

	"github.com/JacobRWebb/InventoryManagement/pkg/middlewares"
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
	user, _ := middlewares.GetUserFromSession(r)

	Render(w, r, pages.Index(user))
}
