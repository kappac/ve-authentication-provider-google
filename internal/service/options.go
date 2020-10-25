package service

// NewOption is a builder for constructor options.
type NewOption func(*veAuthenticationProviderGoogle)

// WithAddress sets address server will listen to.
func WithAddress(a string) NewOption {
	return func(p *veAuthenticationProviderGoogle) {
		p.address = a
	}
}
