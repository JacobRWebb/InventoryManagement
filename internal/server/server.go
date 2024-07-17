package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/JacobRWebb/InventoryManagement/internal/config"
	"github.com/JacobRWebb/InventoryManagement/internal/util/cert"
	"github.com/go-chi/chi/v5"
	"golang.org/x/net/http2"
)

func NewServer(cfg *config.Config, r *chi.Mux) (*http.Server, error) {
	tlsConfig, err := cert.NewTLSConfig(cfg)

	if err != nil {
		return nil, err
	}

	srv := &http.Server{
		Addr:      fmt.Sprintf(":%d", cfg.ApplicationConfig.Port),
		Handler:   r,
		TLSConfig: tlsConfig,
	}

	return srv, nil
}

func Run(srv *http.Server) {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		err := listen(srv)

		if err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}

func listen(srv *http.Server) error {
	http2.ConfigureServer(srv, &http2.Server{})

	return srv.ListenAndServeTLS("", "")
}
