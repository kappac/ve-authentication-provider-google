package middleware

import (
	"github.com/kappac/ve-authentication-provider-google/internal/logger"
	"github.com/kappac/ve-authentication-provider-google/internal/service"
)

// LoggingOption is an option updater for a Logging constructor
type LoggingOption func(*loggingMiddleware)

// WithService sets a services references to wrap with.
func WithService(s service.VEAuthenticationProviderGoogle) LoggingOption {
	return func(l *loggingMiddleware) {
		l.next = s
	}
}

// WithLogger sets a logger to use for logging.
func WithLogger(logger logger.Logger) LoggingOption {
	return func(l *loggingMiddleware) {
		l.logger = logger
	}
}
