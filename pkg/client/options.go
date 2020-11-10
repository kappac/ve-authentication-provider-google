package client

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
