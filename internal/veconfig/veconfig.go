package veconfig

import (
	hc "github.com/kappac/restx-helpers/pkg/config"
)

const (
	envPrefix          string = "VE_GLOBAL_"
	configDefaultDebug bool   = true
)

func getDefaultConfig() *Config {
	return &Config{
		Debug: configDefaultDebug,
	}
}

// Config describes configuration file structure.
type Config struct {
	Debug bool `json:"DEBUG"`
}

var conf *Config

func init() {
	conf = getDefaultConfig()
	config := hc.New(
		hc.WithStrippedPrefix(envPrefix),
	)

	config.Scan(conf)
}

// GetConfig returns current config for a service
func GetConfig() *Config {
	return conf
}
