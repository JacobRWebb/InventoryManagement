package middlewares

import (
	"fmt"
	"net/http"

	"github.com/JacobRWebb/InventoryManagement/pkg/models"
	"github.com/JacobRWebb/InventoryManagement/pkg/store"
	"github.com/gorilla/sessions"
)

type userMiddleware struct {
	store *store.Store
}

func NewUserMiddleware(store *store.Store) UserMiddleware {
	um := &userMiddleware{
		store: store,
	}

	return um
}

func (um userMiddleware) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, exist := GetAuthFromSession(r)

		if !exist {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func getSession(r *http.Request) *sessions.Session {
	fmt.Printf("%v", SessionStore)
	session, err := SessionStore.Get(r, "InventoryManagement")
	fmt.Printf(err.Error())

	return session
}

func SaveAuthToSession(authResponse *models.AuthResponse, w http.ResponseWriter, r *http.Request) {
	userSession := getSession(r)
	fmt.Println(userSession)
	userSession.Values["auth_response"] = *authResponse
	userSession.Save(r, w)
}

func GetAuthFromSession(r *http.Request) (*models.AuthResponse, bool) {
	userSession := getSession(r)
	user, ok := userSession.Values["auth_response"].(*models.AuthResponse)

	if !ok {
		return nil, false
	}

	return user, true
}

func UserLogoutFromSession(w http.ResponseWriter, r *http.Request) {
	userSession := getSession(r)

	userSession.Values["auth_response"] = nil
	userSession.Options.MaxAge = -1
	userSession.Save(r, w)
}
