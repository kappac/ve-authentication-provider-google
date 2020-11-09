package connectionpool

// PoolOption is used to initialize a ConnectionPool
type PoolOption func(*connectionPool)

// WithMin sets a minimum amount of connections ready
// to serve requests at any time.
// Default value is 2.
func WithMin(m int) PoolOption {
	return func(p *connectionPool) {
		p.min = m
	}
}

// WithMax sets a maximum amount of connections ready
// to serve requests at any time. ConnectionPool starts
// with min amount of connections, creating a new connection
// if there is no free connection to serve a request.
// There gould be only max amount of connections per
// ConnectionPool.
// Default value is 5.
func WithMax(m int) PoolOption {
	return func(p *connectionPool) {
		p.max = m
	}
}

// WithConstructor sets a constructor used to initializa
// new connection.
func WithConstructor(c ConnectionConstructor) PoolOption {
	return func(p *connectionPool) {
		p.constructor = c
	}
}
