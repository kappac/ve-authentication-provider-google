package middleware

import (
	"time"

	"github.com/go-kit/kit/log"
	"github.com/kappac/ve-authentication-provider-google/internal/google"
	"github.com/kappac/ve-authentication-provider-google/internal/server"
)

type Logging struct {
	server.VEAuthenticationProviderGoogle
	log.Logger
}

func (m Logging) ValidateToken(token string) (*google.Token, error) {
	defer func(begin time.Time) {
		_ = m.Logger.Log(
			"method", "ValidateToken",
			"token", token,
			"took", time.Since(begin),
		)
	}(time.Now())
	info, err := m.VEAuthenticationProviderGoogle.ValidateToken(token)
	return info, err
}
