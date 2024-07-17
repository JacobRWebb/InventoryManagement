package cert

import (
	"crypto/tls"
	"fmt"

	"github.com/JacobRWebb/InventoryManagement/internal/config"
)

func NewTLSConfig(cfg *config.Config) (*tls.Config, error) {
	certificate, err := tls.LoadX509KeyPair(cfg.CertsConfig.CertificateLocation, cfg.CertsConfig.KeyLocation)

	if err != nil {
		return nil, fmt.Errorf("error while loading certs: %v", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{certificate},
	}

	return tlsConfig, nil
}
