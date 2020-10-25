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

// New constructs a VEAuthenticationProviderGoogle instance.
func New() VEAuthenticationProviderGoogle {
	tv := google.NewTokenVerifier()

	return &veAuthenticationProviderGoogle{
		tv: tv,
	}
}

func (s *veAuthenticationProviderGoogle) ValidateToken(t string) (*google.Token, error) {
	return s.tv.Verify(t)
}

func (s *veAuthenticationProviderGoogle) Run() error {
	var ch = make(chan bool)

	go (func() {
		s.tv.Run()
		ch <- true
	})()

	<-ch

	return nil
}

func (s *veAuthenticationProviderGoogle) Stop() error {
	return s.tv.Stop()
}
