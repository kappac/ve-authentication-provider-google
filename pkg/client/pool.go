package client

import (
	"context"

	"github.com/kappac/ve-authentication-provider-google/internal/connectionpool"
	"github.com/kappac/ve-authentication-provider-google/internal/grpcclient"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/providerinfo"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/request"
	"google.golang.org/grpc"
)

type veAuthenticationProviderPool struct {
	pool     connectionpool.ConnectionPool
	min, max int
}

// New instantiates new Client
func New(opts ...Option) VEAuthenticationProviderGoogleClient {
	p := &veAuthenticationProviderPool{
		min: 2,
		max: 5,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func (p *veAuthenticationProviderPool) Dial(addr string, opts ...grpc.DialOption) error {
	p.pool = connectionpool.New(
		connectionpool.WithMin(p.min),
		connectionpool.WithMax(p.max),
		connectionpool.WithConstructor(func() (grpcclient.Closer, error) {
			return createConnection(addr, opts)
		}),
	)
	return p.pool.Run()
}

func (p *veAuthenticationProviderPool) Close() error {
	return p.pool.Stop()
}

func (p *veAuthenticationProviderPool) ValidateToken(c context.Context, r request.VEValidateTokenRequest) (providerinfo.VEProviderInfo, error) {
	con := p.pool.Pop().(VEAuthenticationProviderGoogleClient)
	defer p.pool.Push(con)
	return con.ValidateToken(c, r)
}

func createConnection(addr string, opts []grpc.DialOption) (VEAuthenticationProviderGoogleClient, error) {
	var err error

	client := newClient()
	if err = client.Dial(addr, opts...); err != nil {
		return nil, err
	}

	return client, err
}
