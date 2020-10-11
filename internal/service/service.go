package service

import (
	"github.com/micro/micro/v3/service"

	"github.com/kappac/ve-authentication-provider-google/internal/api"
	"github.com/kappac/ve-authentication-provider-google/internal/config"
	"github.com/kappac/ve-authentication-provider-google/internal/handler"
)

var (
	srv *service.Service
)

// GetService - singleton, to get a service instance
func GetService() *service.Service {
	if srv != nil {
		return srv
	}

	cfg := config.GetConfig()

	srv = service.New(
		service.Name(cfg.ServiceName),
	)

	srv.Init()

	api.RegisterRestXLoggingServiceHandler(srv.Server(), new(handler.Handler))

	return srv
}
