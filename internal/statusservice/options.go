package statusservice

// NewOption is an option updater for the constructor
type NewOption func(s *statusService)

// WithAddress sets an address service will listen at.
func WithAddress(a string) NewOption {
	return func(s *statusService) {
		s.address = a
	}
}

// WithEndpointSource sets a source for an apropriate Probe.
func WithEndpointSource(p Probe, src SourceSubscriber) NewOption {
	return func(s *statusService) {
		s.sources[p] = src
	}
}
