package google

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/micro/micro/v3/service/logger"
)

// TokenVerifier is JWT token verification
// entity for Google asymetric sign in.
type TokenVerifier struct {
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
func NewTokenVerifier() *TokenVerifier {
	certs := newOauthCertificates()
	closeCh := make(chan chan error)

	return &TokenVerifier{
		certs:   certs,
		closeCh: closeCh,
	}
}

// Run starts TokenVerifier execution loop
func (tv *TokenVerifier) Run() {
	for !tv.isClosing {
		select {
		case errc := <-tv.closeCh:
			logger.Debug("Closing")
			tv.isClosing = true
			tv.err = tv.certs.stop()
			tv.closeChannels()
			errc <- tv.err
		}
	}
}

// Stop stops TokenVerifier execution loop
func (tv *TokenVerifier) Stop() error {
	cc := make(chan error)

	tv.closeCh <- cc

	select {
	case errc := <-cc:
		return errc
	}
}

// Verify validates token.
func (tv *TokenVerifier) Verify(t string) (*Token, error) {
	token, err := jwt.ParseWithClaims(t, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		logger.Debugf("parsed token: %v\n", token)

		kid := token.Header["kid"].(string)

		cert, err := tv.certs.get(kid)
		if err != nil {
			return nil, err
		}

		return cert, nil
	})

	if err != nil {
		return nil, err
	}

	return tv.mapToken(token), nil
}

func (tv *TokenVerifier) mapToken(t *jwt.Token) *Token {
	claims := t.Claims.(*tokenClaims)

	return &Token{
		FullName:   claims.FullName,
		GivenName:  claims.GivenName,
		FamilyName: claims.FamilyName,
		Email:      claims.Email,
		Picture:    claims.Picture,
	}
}

func (tv *TokenVerifier) closeChannels() {
	close(tv.closeCh)
}
