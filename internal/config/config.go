package config

import (
	"github.com/kappac/ve-authentication-provider-google/internal/veconfig"
)

// ServiceConfig stores parameters passed to the service.
type ServiceConfig struct {
	veconfig.DefaultServiceParameters
}

// Verify checks if all the parameters are valid.
func (p ServiceConfig) Verify() error {
	return p.DefaultServiceParameters.Verify()
}

const (
	defaultName          = "ve-authentication-provider-google"
	defaultAddress       = ":8000"
	defaultProbesAddress = ":8010"
)

var (
	// Config is a set of service parameters.
	Config ServiceConfig = ServiceConfig{
		DefaultServiceParameters: veconfig.DefaultServiceParameters{
			Name:          defaultName,
			Address:       defaultAddress,
			ProbesAddress: defaultProbesAddress,
		},
	}
)

func init() {
	veconfig.Parse(&Config)
}
