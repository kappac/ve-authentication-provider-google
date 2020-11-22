package client

import (
	"context"

	"github.com/kappac/ve-authentication-provider-google/pkg/proto/providerinfo"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/request"
	"github.com/kappac/ve-back-end-utils/pkg/connectionpool"
	"github.com/kappac/ve-back-end-utils/pkg/service"
	"google.golang.org/grpc"
)

type veAuthenticationProviderPool struct {
	pool     connectionpool.ConnectionPool
	min, max int
	addr     string
	grpcOpts []grpc.DialOption
}

// New instantiates new Client
func New(opts ...Option) VEAuthenticationProviderGoogleClient {
	p := &veAuthenticationProviderPool{
		min:      2,
		max:      5,
		grpcOpts: make([]grpc.DialOption, 0),
	}

	for _, opt := range opts {
		opt(p)
	}

	p.pool = connectionpool.New(
		connectionpool.WithMin(p.min),
		connectionpool.WithMax(p.max),
		connectionpool.WithConstructor(p.createConnection),
	)

	return p
}

func (p *veAuthenticationProviderPool) Dial() error {
	return p.pool.Run()
}

func (p *veAuthenticationProviderPool) Close() error {
	return p.pool.Stop()
}

func (p *veAuthenticationProviderPool) ValidateToken(c context.Context, r request.VEValidateTokenRequest) (providerinfo.VEProviderInfo, error) {
	con, _ := p.pool.Get(context.Background())
	defer p.pool.Put(con)
	coni := con.(VEAuthenticationProviderGoogleClient)
	return coni.ValidateToken(c, r)
}

func (p *veAuthenticationProviderPool) createConnection() (service.Closer, error) {
	var err error

	client := newClient(p.addr, p.grpcOpts)
	if err = client.Dial(); err != nil {
		return nil, err
	}

	return client, err
}
