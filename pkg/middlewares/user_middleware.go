package middlewares

import (
	"net/http"

	"github.com/JacobRWebb/InventoryManagement/pkg/store"
	"github.com/gorilla/sessions"
)

var SessionStore = sessions.NewCookieStore([]byte("Secret-Key"))

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
		// session, _ := SessionStore.Get(r, "session-name")
		// user, ok := session.Values["user"].(models.User)

		// if !ok || user != models.User{} {
		// 	if r.URL.Path != "/login" && r.URL.Path != "/register" {
		// 		http.Redirect(w, r, "/login", http.StatusSeeOther)
		// 		return
		// 	}
		// }

		next.ServeHTTP(w, r)
	}
}
