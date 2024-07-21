package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/JacobRWebb/InventoryManagement/internal/models"
	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

const (
	SessionName = "Inventory-Management"
	AuthKey     = "auth_response"
)

func init() {
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
}

func GetSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, SessionName)
}

func SetAuthResponse(w http.ResponseWriter, r *http.Request, auth *models.AuthResponse) error {
	session, err := store.Get(r, SessionName)
	if err != nil {
		return err
	}

	authJSON, err := json.Marshal(auth)
	if err != nil {
		return err
	}

	session.Values[AuthKey] = string(authJSON)
	return session.Save(r, w)
}

func GetAuthResponse(r *http.Request) (*models.AuthResponse, error) {
	session, err := store.Get(r, SessionName)
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %w", err)
	}

	authValue, exists := session.Values[AuthKey]

	if authValue == nil || !exists {
		return nil, errors.New("AuthKey not found in session")
	}

	// Type assertion
	authJSON, ok := authValue.(string)
	if !ok {
		return nil, errors.New("AuthKey value is not a string")
	}

	if authJSON == "" {
		return nil, errors.New("AuthKey value is empty")
	}

	var auth models.AuthResponse
	err = json.Unmarshal([]byte(authJSON), &auth)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal auth JSON: %w", err)
	}

	return &auth, nil
}

func ClearAuthResponse(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, SessionName)
	if err != nil {
		return err
	}

	delete(session.Values, AuthKey)
	return session.Save(r, w)
}

func ClearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, SessionName)
	if err != nil {
		return err
	}
	session.Values = make(map[interface{}]interface{})
	return session.Save(r, w)
}
