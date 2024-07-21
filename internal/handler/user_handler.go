package handler

import (
	"net/http"

	"github.com/JacobRWebb/InventoryManagement/internal/models"
	"github.com/JacobRWebb/InventoryManagement/internal/service"
	session "github.com/JacobRWebb/InventoryManagement/internal/util"
	"github.com/JacobRWebb/InventoryManagement/internal/web/templates/forms"
	"github.com/JacobRWebb/InventoryManagement/internal/web/templates/pages"
	"google.golang.org/grpc/status"
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

	err := h.services.UserService.RegisterUser(values.Email, values.Password)

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

	authRes, err := h.services.UserService.LoginUser(values.Email, values.Password)

	if err != nil {
		if grpcErr, ok := status.FromError(err); ok {
			errors["form"] = grpcErr.Message()
		} else {
			errors["form"] = err.Error()
		}

		Render(w, r, forms.LoginAccountForm(values, errors))
		return
	}

	err = session.SetAuthResponse(w, r, authRes)

	if err != nil {

		errors["form"] = "Interal server error."

		Render(w, r, forms.LoginAccountForm(values, errors))
		return
	}

	HxRedirect(w, r, "/")
}

func (h *userHandler) HandleUserLogoutGet(w http.ResponseWriter, r *http.Request) {
	session.ClearAuthResponse(w, r)

	HxRedirect(w, r, "/login")
}

func (h *userHandler) HandleUserLogoutPost(w http.ResponseWriter, r *http.Request) {
	session.ClearAuthResponse(w, r)

	HxRedirect(w, r, "/login")
}
