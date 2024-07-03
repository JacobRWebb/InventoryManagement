package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

var (
	workDir, _ = os.Getwd()
)

type basicHandler struct {
}

func NewBasicHandler() BasicHandler {
	bh := &basicHandler{}
	return bh
}

func (bh *basicHandler) HandleStatic(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filesDir := http.Dir(filepath.Join(workDir, "/pkg/web/static"))
		fs := http.FileServer(filesDir)
		http.StripPrefix("/static/", fs).ServeHTTP(w, r)
	}
}

func (bh *basicHandler) HandleFaviconGet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(workDir, "/pkg/web/static/assets/favicon.ico"))
}
