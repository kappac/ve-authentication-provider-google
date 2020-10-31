package main

import (
	"github.com/kappac/ve-authentication-provider-google/internal/config"
	"github.com/kappac/ve-authentication-provider-google/internal/launch"
	"github.com/kappac/ve-authentication-provider-google/internal/logger"
	"github.com/kappac/ve-authentication-provider-google/internal/service"
	"github.com/kappac/ve-authentication-provider-google/internal/statusservice"
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

	launchErr := <-launch.Launch(
		launch.WithRunStopper(svc),
		launch.WithRunStopper(ssvc),
	)

	log.Infom("Terminating", "err", launchErr)
}
