package config

import (
	hc "github.com/kappac/restx-helpers/pkg/config"
	"github.com/micro/micro/v3/service/logger"
)

const (
	envPrefix                string = "VE_"
	configDefaultServiceName string = "com.venueexplorer.authorization.provider.google"
)

func getDefaultConfig() *Config {
	return &Config{
		ServiceName: configDefaultServiceName,
	}
}

// Config describes configuration file structure.
type Config struct {
	ServiceName string `json:"SERVICE_NAME"`
}

var conf *Config

func init() {
	conf = getDefaultConfig()
	config := hc.New(
		hc.WithStrippedPrefix(envPrefix),
	)

	config.Scan(conf)

	logger.Debugf("Config: %#v", conf)
}

// GetConfig returns current config for a service
func GetConfig() *Config {
	return conf
}
