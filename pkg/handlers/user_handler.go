package handlers

import (
	"net/http"

	"github.com/JacobRWebb/InventoryManagement/pkg/middlewares"
	"github.com/JacobRWebb/InventoryManagement/pkg/models"
	"github.com/JacobRWebb/InventoryManagement/pkg/store"
	"github.com/JacobRWebb/InventoryManagement/pkg/web/templates/forms"
	"github.com/JacobRWebb/InventoryManagement/pkg/web/templates/pages"
)

type userHandler struct {
	store *store.Store
}

func NewUserHandler(store *store.Store) UserHandler {
	uh := &userHandler{
		store: store,
	}

	return uh
}

func (h userHandler) HandleUserCreateGet(w http.ResponseWriter, r *http.Request) {
	values, _ := models.ParseCreateAccountFormValuesAndValidate(r)

	Render(w, r, pages.Register(values))
}

func (h userHandler) HandleUserCreatePost(w http.ResponseWriter, r *http.Request) {
	values, errors := models.ParseCreateAccountFormValuesAndValidate(r)

	if len(errors) > 0 {
		Render(w, r, forms.CreateAccountForm(values, errors))
		return
	}

	// TODO Create Account

	HxRedirect(w, r, "/login")
}

func (h *userHandler) HandleUserLoginGet(w http.ResponseWriter, r *http.Request) {
	values, _ := models.ParseLoginAccountFormValuesAndValidate(r)

	Render(w, r, pages.Login(values))
}

func (h *userHandler) HandleUserLoginPost(w http.ResponseWriter, r *http.Request) {
	values, errors := models.ParseLoginAccountFormValuesAndValidate(r)

	if len(errors) > 0 {
		Render(w, r, forms.LoginAccountForm(values, errors))
		return
	}

	// TODO Login User
	user := &models.User{
		Username: values.Email,
		Profile: &models.UserProfile{
			ProfilePic: "https://ui-avatars.com/api/?name=Jacob+Webb",
			Email:      values.Email,
		},
	}

	session, _ := middlewares.SessionStore.Get(r, "session-name")
	session.Values["user"] = user
	session.Save(r, w)

	HxRedirect(w, r, "/")
}

func (h *userHandler) HandleUserLogoutPost(w http.ResponseWriter, r *http.Request) {
	// TODO logout user session.

	HxRedirect(w, r, "/")
}
