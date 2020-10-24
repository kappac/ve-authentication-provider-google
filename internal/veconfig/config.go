package veconfig

import (
	"github.com/kappac/restx-helpers/pkg/config"
)

const (
	serviceEnvPrefix = "VE_SERVICE_"
)

// Parse populates "p" with values provided as
// service parameters. The parameters are read
// from env variables prefixed with "VE_SERVICE_".
func Parse(p ServiceParameters) error {
	c := config.New(
		config.WithStrippedPrefix(serviceEnvPrefix),
	)
	return c.Scan(p)
}

// ParsePrefix populates "p" with values provided as
// service parameters. The parameters are read
// from env variables prefixed with "ep".
func ParsePrefix(ep string, p ServiceParameters) error {
	c := config.New(
		config.WithStrippedPrefix(ep),
	)
	return c.Scan(p)
}
