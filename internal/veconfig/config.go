package veconfig

import (
	"github.com/kappac/restx-helpers/pkg/config"
)

const (
	serviceEnvPrefix = "VE_SERVICE_"
)

type parseOptions struct {
	envPrefix string
}

// ParseOption is an option updater for Parse function.
type ParseOption func(o *parseOptions)

// Parse populates "p" with values provided as
// service parameters. The parameters are read
// from env variables.
// Default env prefix is "VE_SERVICE_".
func Parse(p ServiceParameters, os ...ParseOption) error {
	options := &parseOptions{
		envPrefix: serviceEnvPrefix,
	}

	for _, o := range os {
		o(options)
	}

	c := config.New(
		config.WithStrippedPrefix(options.envPrefix),
	)

	return c.Scan(p)
}

// WithEnvPrefix sets a prefix to be checked against
// env variables.
func WithEnvPrefix(p string) ParseOption {
	return func(o *parseOptions) {
		o.envPrefix = p
	}
}
