package service

import (
	"net"

	"github.com/kappac/ve-authentication-provider-google/internal/logger"
	"github.com/kappac/ve-authentication-provider-google/internal/pb"
	"github.com/kappac/ve-authentication-provider-google/internal/types/runstopper"
	"google.golang.org/grpc"
)

// VEAuthenticationProviderGoogle manages service API.
type VEAuthenticationProviderGoogle interface {
	runstopper.RunStopper
}

type veAuthenticationProviderGoogle struct {
	address    string
	binding    grpcBinding
	grpcserver *grpc.Server
	logger     logger.Logger
}

// New constructs a VEAuthenticationProviderGoogle instance
func New(os ...NewOption) VEAuthenticationProviderGoogle {
	b := grpcBinding{
		svc: newAuthProviderGoogle(),
	}
	l := logger.New(
		logger.WithEntity("VEAuthenticationProviderGoogle"),
	)
	p := &veAuthenticationProviderGoogle{
		binding: b,
		logger:  l,
	}

	for _, o := range os {
		o(p)
	}

	return p
}

func (p *veAuthenticationProviderGoogle) Run() error {
	var errc = make(chan error)

	go p.runBinding(errc)
	go p.runGrpc(errc)

	select {
	case err := <-errc:
		return err
	}
}

func (p *veAuthenticationProviderGoogle) Stop() error {
	p.grpcserver.Stop()
	return p.binding.Stop()
}

func (p *veAuthenticationProviderGoogle) runBinding(errc chan<- error) {
	errc <- p.binding.Run()
}

func (p *veAuthenticationProviderGoogle) runGrpc(errc chan<- error) {
	ln, err := net.Listen("tcp", p.address)
	if err != nil {
		errc <- err
		return
	}

	p.grpcserver = grpc.NewServer()

	pb.RegisterVEAuthProviderGoogleServiceServer(p.grpcserver, p.binding)

	errc <- p.grpcserver.Serve(ln)
}