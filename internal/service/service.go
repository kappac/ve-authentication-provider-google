package service

import (
	"github.com/kappac/ve-authentication-provider-google/internal/google"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/providerinfo"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/request"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/response"
	veerror "github.com/kappac/ve-back-end-utils/pkg/proto/error"
	veservice "github.com/kappac/ve-back-end-utils/pkg/service"
	"github.com/kappac/ve-back-end-utils/pkg/statusservice"
)

type authProviderGoogle interface {
	veservice.RunStopper
	ValidateToken(req request.VEValidateTokenRequest) (response.VEValidateTokenResponse, error)
	GetStatisticsSource() statusservice.SourceSubscriber
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
	c := make(chan error)

	go (func() {
		p.tv.Run()
		c <- nil
	})()

	return <-c
}

func (p *authProviderGoogleImpl) Stop() error {
	return p.tv.Stop()
}

func (p *authProviderGoogleImpl) GetStatisticsSource() statusservice.SourceSubscriber {
	return google.NewStatisticsSource(google.WithTokenVerifier(p.tv))
}
