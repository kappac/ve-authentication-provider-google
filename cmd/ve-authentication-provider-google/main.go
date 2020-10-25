package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kappac/ve-authentication-provider-google/internal/config"
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
	errc := make(chan error)

	go func() {
		errc <- interrupt()
	}()

	go func() {
		ln, err := net.Listen("tcp", config.Config.Address)
		if err != nil {
			errc <- err
			return
		}
		s := grpc.NewServer()
		pb.RegisterVEAuthProviderGoogleServiceServer(s, service.GrpcBinding{svc})
		errc <- s.Serve(ln)
	}()

	_ = log.Info("fatal", <-errc)
}

func interrupt() error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return fmt.Errorf("%s", <-c)
}
