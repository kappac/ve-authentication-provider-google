package server

import (
	"github.com/kappac/ve-authentication-provider-google/internal/google"
)

// VEAuthenticationProviderGoogle represents service API
type VEAuthenticationProviderGoogle interface {
	ValidateToken(t string) (*google.Token, error)
}
