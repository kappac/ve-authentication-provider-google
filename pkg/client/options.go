package client

import "google.golang.org/grpc"

// Option is used to configure Client instance
type Option func(*veAuthenticationProviderPool)

// WithMinConnections sets minimum connections available
// to Client.
// Default value: 2
func WithMinConnections(m int) Option {
	return func(p *veAuthenticationProviderPool) {
		p.min = m
	}
}

// WithMaxConnections sets maximum connections available
// to Client.
// Default value: 5
func WithMaxConnections(m int) Option {
	return func(p *veAuthenticationProviderPool) {
		p.max = m
	}
}

// WithAddress sets an address client will connect to
func WithAddress(a string) Option {
	return func(p *veAuthenticationProviderPool) {
		p.addr = a
	}
}

// WithGRPCDialOptions sets GRPC client will use to dial
// to a server
func WithGRPCDialOptions(opts ...grpc.DialOption) Option {
	return func(p *veAuthenticationProviderPool) {
		p.grpcOpts = append(p.grpcOpts, opts...)
	}
}
