package middleware

import (
	"time"

	"github.com/kappac/ve-authentication-provider-google/internal/google"
	"github.com/kappac/ve-authentication-provider-google/internal/logger"
	"github.com/kappac/ve-authentication-provider-google/internal/service"
)

// LoggingMiddleware is a logging middleware
type LoggingMiddleware interface {
	ValidateToken(token string) (*google.Token, error)
}

type loggingMiddleware struct {
	next   service.VEAuthenticationProviderGoogle
	logger logger.Logger
}

// NewLogging is a Logging constructor
func NewLogging(os ...LoggingOption) LoggingMiddleware {
	l := &loggingMiddleware{
		logger: logger.New(),
	}

	for _, o := range os {
		o(l)
	}

	return l
}

// ValidateToken is a proxy for service ValidateToken for logging.
func (m *loggingMiddleware) ValidateToken(token string) (*google.Token, error) {
	if m.next == nil {
		return nil, errorLoggingNoService
	}

	defer func(begin time.Time) {
		_ = m.logger.Info(
			"method", "ValidateToken",
			"token", token,
			"took", time.Since(begin),
		)
	}(time.Now())

	info, err := m.next.ValidateToken(token)

	return info, err
}
