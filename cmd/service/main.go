package main

import (
	"math/rand"
	"time"

	"github.com/kappac/ve-authentication-provider-google/internal/config"
	"github.com/kappac/ve-authentication-provider-google/internal/launch"
	"github.com/kappac/ve-authentication-provider-google/internal/logger"
	"github.com/kappac/ve-authentication-provider-google/internal/runstopper"
	"github.com/kappac/ve-authentication-provider-google/internal/service"
)

func main() {
	log := logger.New(
		logger.WithEntity("Service"),
	)
	svc := service.New(
		service.WithAddress(config.Config.Address),
	)

	rand.Seed(time.Now().UnixNano())

	launchErr := <-launch.Launch(
		svcLauncher(svc),
	)

	log.Infom("Terminating", "err", launchErr)
}

func svcLauncher(svc runstopper.RunStopper) launch.Function {
	return func(errCh chan<- error, exitCh chan bool) {
		var (
			isClosing bool
		)

		go (func() {
			errCh <- svc.Run()
		})()

		for !isClosing {
			select {
			case <-exitCh:
				isClosing = true
				_ = svc.Stop()
				exitCh <- true
			}
		}
	}
}
