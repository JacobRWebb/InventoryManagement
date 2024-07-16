package handlers

import (
	"fmt"
	"net/http"

	"github.com/JacobRWebb/InventoryManagement/pkg/middlewares"
	"github.com/JacobRWebb/InventoryManagement/pkg/models"
	"github.com/JacobRWebb/InventoryManagement/pkg/store"
	"github.com/JacobRWebb/InventoryManagement/pkg/web/templates/forms"
	"github.com/JacobRWebb/InventoryManagement/pkg/web/templates/pages"
	"google.golang.org/grpc/status"
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

	_, err := h.store.UserStore.RegisterUser(values.Email, values.Password)

	if err != nil {
		if grpcErr, ok := status.FromError(err); ok {
			errors["form"] = grpcErr.Message()
		} else {
			errors["form"] = err.Error()
		}

		Render(w, r, forms.CreateAccountForm(values, errors))

		return
	}

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

	ar, err := h.store.UserStore.LoginUser(values.Email, values.Password)

	if err != nil {
		if grpcErr, ok := status.FromError(err); ok {
			errors["form"] = grpcErr.Message()
		} else {
			errors["form"] = err.Error()
		}

		Render(w, r, forms.LoginAccountForm(values, errors))
		return
	}

	fmt.Printf("\nHandlerUserLoginPost %v", ar)

	middlewares.SaveAuthToSession(ar, w, r)

	HxRedirect(w, r, "/")
}

func (h *userHandler) HandleUserLogoutPost(w http.ResponseWriter, r *http.Request) {
	middlewares.UserLogoutFromSession(w, r)
	HxRedirect(w, r, "/")
}
