package service

import (
	"github.com/kappac/ve-authentication-provider-google/internal/google"
	veerror "github.com/kappac/ve-authentication-provider-google/internal/types/error"
	"github.com/kappac/ve-authentication-provider-google/internal/types/providerinfo"
	"github.com/kappac/ve-authentication-provider-google/internal/types/request"
	"github.com/kappac/ve-authentication-provider-google/internal/types/response"
	"github.com/kappac/ve-authentication-provider-google/internal/types/runstopper"
)

type authProviderGoogle interface {
	runstopper.RunStopper
	ValidateToken(req request.VEValidateTokenRequest) (response.VEValidateTokenResponse, error)
}

type authProviderGoogleImpl struct {
	tv google.TokenVerifier
}

func newAuthProviderGoogle() authProviderGoogle {
	tv := google.NewTokenVerifier()
	p := &authProviderGoogleImpl{
		tv: tv,
	}
	return p
}

func (p *authProviderGoogleImpl) ValidateToken(req request.VEValidateTokenRequest) (response.VEValidateTokenResponse, error) {
	var (
		veinfo providerinfo.VEProviderInfo
		veerr  veerror.VEError
	)

	if t, err := p.tv.Verify(req.GetToken()); err == nil {
		veinfo = providerinfo.New(
			providerinfo.WithFullName(t.FullName),
			providerinfo.WithGivenName(t.GivenName),
			providerinfo.WithFamilyName(t.FamilyName),
			providerinfo.WithEmail(t.Email),
			providerinfo.WithPicture(t.Picture),
		)
	} else {
		if e, ok := err.(veerror.VEError); ok {
			veerr = e
		}

		veerr = veerror.New(
			veerror.WithDescription(err.Error()),
		)
	}

	resp := response.New(
		response.WithRequest(req),
		response.WithInfo(veinfo),
		response.WithError(veerr),
	)

	return resp, nil
}

func (p *authProviderGoogleImpl) Run() error {
	c := make(chan bool)

	go (func() {
		p.tv.Run()
		c <- true
	})()

	select {
	case <-c:
		return nil
	}
}

func (p *authProviderGoogleImpl) Stop() error {
	return p.tv.Stop()
}
