package main

import (
	"flag"
	"fmt"
	stdlog "log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/kappac/ve-authentication-provider-google/internal/middleware"
	"github.com/kappac/ve-authentication-provider-google/internal/pb"
	"github.com/kappac/ve-authentication-provider-google/internal/server"
	"github.com/kappac/ve-authentication-provider-google/internal/service"
	"google.golang.org/grpc"
)

func main() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		addr = fs.String("addr", ":8002", "Address for gRPC server")
	)
	flag.Usage = fs.Usage // only show our flags
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	// package metrics
	// var requestDuration metrics.TimeHistogram
	// {
	// 	requestDuration = metrics.NewTimeHistogram(time.Nanosecond, metrics.NewMultiHistogram(
	// 		expvar.NewHistogram("request_duration_ns", 0, 5e9, 1, 50, 95, 99),
	// 		prometheus.NewSummary(stdprometheus.SummaryOpts{
	// 			Namespace: "myorg",
	// 			Subsystem: "addsvc",
	// 			Name:      "duration_ns",
	// 			Help:      "Request duration in nanoseconds.",
	// 		}, []string{"method"}),
	// 	))
	// }

	// package log
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		// logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC).With("caller", log.DefaultCaller)
		stdlog.SetFlags(0)                             // flags are handled by Go kit's logger
		stdlog.SetOutput(log.NewStdlibAdapter(logger)) // redirect anything using stdlib log to us
	}

	// Business domain
	var svc server.VEAuthenticationProviderGoogle
	{
		svc = service.NewService()
		svc = middleware.Logging{svc, logger}
		// svc = middleware.Instrumenting{svc, requestDuration}
	}

	// Mechanical stuff
	rand.Seed(time.Now().UnixNano())
	// root := context.Background()
	errc := make(chan error)

	go func() {
		errc <- interrupt()
	}()

	go func() {
		// transportLogger := log.NewContext(logger).With("transport", "gRPC")
		ln, err := net.Listen("tcp", *addr)
		if err != nil {
			errc <- err
			return
		}
		s := grpc.NewServer() // uses its own, internal context
		pb.RegisterVEAuthProviderGoogleServiceServer(s, service.GrpcBinding{svc})
		// _ = transportLogger.Log("addr", *addr)
		errc <- s.Serve(ln)
	}()

	_ = logger.Log("fatal", <-errc)
}

func interrupt() error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return fmt.Errorf("%s", <-c)
}
