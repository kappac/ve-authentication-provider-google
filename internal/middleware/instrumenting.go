package middleware

// import (
// 	"time"

// 	"github.com/go-kit/kit/metrics"
// 	"github.com/kappac/ve-authentication-provider-google/internal/google"
// 	"github.com/kappac/ve-authentication-provider-google/internal/server"
// )

// type Instrumenting struct {
// 	server.VEAuthenticationProviderGoogle
// 	requestDuration metrics.TimeHistogram
// }

// func (m Instrumenting) ValidateToken(token string) (*google.Token, error) {
// 	defer func(begin time.Time) {
// 		methodField := metrics.Field{Key: "method", Value: "VelidateToken"}
// 		m.requestDuration.With(methodField).Observe(time.Since(begin))
// 	}(time.Now())
// 	info, err := m.VEAuthenticationProviderGoogle.VerifyToken(token)
// 	return info, err
// }
