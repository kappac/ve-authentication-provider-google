package main

import (
	"math/rand"
	"net"
	"time"

	"github.com/kappac/ve-authentication-provider-google/internal/config"
	"github.com/kappac/ve-authentication-provider-google/internal/launch"
	"github.com/kappac/ve-authentication-provider-google/internal/logger"
	"github.com/kappac/ve-authentication-provider-google/internal/middleware"
	"github.com/kappac/ve-authentication-provider-google/internal/pb"
	"github.com/kappac/ve-authentication-provider-google/internal/service"
	"google.golang.org/grpc"
)

func main() {
	log := logger.New(
		logger.WithEntity("Service"),
	)
	svcNoLog := service.New()
	svc := middleware.NewLogging(
		middleware.WithService(svcNoLog),
		middleware.WithLogger(
			logger.New(
				logger.WithEntity("LoggingMiddleware"),
				logger.WithLogger(log),
			),
		),
	)

	rand.Seed(time.Now().UnixNano())

	launchErr := <-launch.Launch(
		svcLauncher(svcNoLog.(launch.RunStopper)),
		grpcBindingLauncher(service.GrpcBinding{svc}),
	)

	_ = log.Infom("Terminating", "err", launchErr)
}

func svcLauncher(svc launch.RunStopper) launch.Function {
	return func(errCh chan<- error, exitCh <-chan bool) {
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
				svc.Stop()
			}
		}
	}
}

func grpcBindingLauncher(srv pb.VEAuthProviderGoogleServiceServer) launch.Function {
	return func(errCh chan<- error, exitCh <-chan bool) {
		var (
			isClosing bool
		)

		ln, err := net.Listen("tcp", config.Config.Address)
		if err != nil {
			errCh <- err
			return
		}

		s := grpc.NewServer()
		pb.RegisterVEAuthProviderGoogleServiceServer(s, srv)

		go (func() {
			errCh <- s.Serve(ln)
		})()

		for !isClosing {
			select {
			case <-exitCh:
				isClosing = true
				s.Stop()
			}
		}
	}
}
