package google

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kappac/ve-authentication-provider-google/internal/logger"
)

// TokenVerifier is JWT token verification
// entity for Google asymetric sign in.
type TokenVerifier interface {
	Run()
	Stop() error
	Verify(t string) (*Token, error)
}

type tokenVerifier struct {
	logger    logger.Logger
	certs     *oauthCertificates
	closeCh   chan chan error
	isClosing bool
	err       error
}

// Token represent token data we are able to get
// from Google's JWT.
type Token struct {
	FullName   string
	GivenName  string
	FamilyName string
	Picture    string
	Email      string
}

// NewTokenVerifier instantiates TokenVerifier
func NewTokenVerifier() TokenVerifier {
	certs := newOauthCertificates()
	closeCh := make(chan chan error)
	l := logger.New(
		logger.WithEntity("TokenVerifier"),
	)

	go certs.run()

	return &tokenVerifier{
		logger:  l,
		certs:   certs,
		closeCh: closeCh,
	}
}

// Run starts TokenVerifier execution loop
func (tv *tokenVerifier) Run() {
	tv.logger.Debugm("starting")

	for !tv.isClosing {
		select {
		case errc := <-tv.closeCh:
			tv.isClosing = true
			tv.err = tv.certs.stop()
			tv.closeChannels()

			tv.logger.Debugm("closing", "err", tv.err)

			errc <- tv.err
		}
	}
}

// Stop stops TokenVerifier execution loop
func (tv *tokenVerifier) Stop() error {
	cc := make(chan error)

	tv.closeCh <- cc

	tv.certs.stop()

	select {
	case errc := <-cc:
		return errc
	}
}

// Verify validates token.
func (tv *tokenVerifier) Verify(t string) (*Token, error) {
	token, err := jwt.ParseWithClaims(t, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		var (
			kid string
		)

		if k, ok := token.Header["kid"]; ok {
			kid = k.(string)
		} else {
			return nil, errorNoCertificateKeyID
		}

		cert, err := tv.certs.get(kid)
		if err != nil {
			return nil, err
		}

		return cert, nil
	})

	if err != nil {
		tv.logger.Infom("Verify", "err", err)
		return nil, err
	}

	return tv.mapToken(token), nil
}

func (tv *tokenVerifier) mapToken(t *jwt.Token) *Token {
	claims := t.Claims.(*tokenClaims)

	return &Token{
		FullName:   claims.FullName,
		GivenName:  claims.GivenName,
		FamilyName: claims.FamilyName,
		Email:      claims.Email,
		Picture:    claims.Picture,
	}
}

func (tv *tokenVerifier) closeChannels() {
	close(tv.closeCh)
}
