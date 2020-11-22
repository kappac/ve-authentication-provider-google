package main

import (
	"github.com/kappac/ve-authentication-provider-google/internal/config"
	"github.com/kappac/ve-authentication-provider-google/internal/service"
	"github.com/kappac/ve-back-end-utils/pkg/launcher"
	"github.com/kappac/ve-back-end-utils/pkg/logger"
	"github.com/kappac/ve-back-end-utils/pkg/statusservice"
)

func main() {
	log := logger.New(
		logger.WithEntity("Service"),
	)
	svc := service.New(
		service.WithAddress(config.Config.Address),
	)
	ssvc := statusservice.New(
		statusservice.WithAddress(config.Config.ProbesAddress),
		statusservice.WithEndpointSource(
			statusservice.ProbeReadiness,
			svc.GetStatisticsSource(),
		),
	)

	l := launcher.New(
		launcher.WithService(svc),
		launcher.WithService(ssvc),
	)
	if err := l.Run(); err != nil {
		log.Infom("Terminating", "err", err)
	}
}
