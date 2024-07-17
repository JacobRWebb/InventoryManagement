package router

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func AttachStaticRoutes(r *chi.Mux) {
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "/internal/web/static"))

	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(filesDir)))
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(workDir, "/internal/web/static/assets/favicon.ico"))
	})
}
