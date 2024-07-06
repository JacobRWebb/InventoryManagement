package middlewares

import (
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
		_, exist := GetUserFromSession(r)

		if !exist {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func getUserSession(r *http.Request) *sessions.Session {
	session, _ := SessionStore.Get(r, "InventoryManagement")
	return session
}

func SaveUserToSession(user *models.SessionUser, w http.ResponseWriter, r *http.Request) {
	userSession := getUserSession(r)
	userSession.Values["user"] = *user
	userSession.Save(r, w)

}

func GetUserFromSession(r *http.Request) (*models.SessionUser, bool) {
	userSession := getUserSession(r)
	user, ok := userSession.Values["user"].(*models.SessionUser)

	if !ok {
		return nil, false
	}

	return user, true
}

func UserLogoutFromSession(w http.ResponseWriter, r *http.Request) {
	userSession := getUserSession(r)

	userSession.Values["user"] = nil
	userSession.Options.MaxAge = -1
	userSession.Save(r, w)
}
