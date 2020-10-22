package service

import (
	"github.com/kappac/ve-authentication-provider-google/internal/google"
)

// VEAuthenticationProviderGoogle represents service API
type VEAuthenticationProviderGoogle interface {
	ValidateToken(t string) (*google.Token, error)
}

type veAuthenticationProviderGoogle struct {
	tv google.TokenVerifier
}

func NewService() *veAuthenticationProviderGoogle {
	tv := google.NewTokenVerifier()

	go tv.Run()

	return &veAuthenticationProviderGoogle{
		tv: tv,
	}
}

func (s *veAuthenticationProviderGoogle) ValidateToken(t string) (*google.Token, error) {
	return s.tv.Verify(t)
}

func (s *veAuthenticationProviderGoogle) Stop() error {
	return s.tv.Stop()
}
