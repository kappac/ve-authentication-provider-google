package google

import (
	"github.com/dgrijalva/jwt-go"
)

const (
	// The value of iss in the ID token is equal to accounts.google.com or https://accounts.google.com.
	tokenVerifierIssuer = "accounts.google.com"
	// The value of aud in the ID token is equal to one of your app's client IDs.
	// This check is necessary to prevent ID tokens issued to a malicious app being
	// used to access data about the same user on your app's backend server.
	tokenVerifierAudience = "needs_to_be_updated_from_env"
)

type tokenClaims struct {
	jwt.StandardClaims

	Email      string `json:"email,omitempty"`
	FullName   string `json:"name,omitempty"`
	GivenName  string `json:"given_name,omitempty"`
	FamilyName string `json:"family_name,omitempty"`
	Picture    string `json:"picture,omitempty"`
}

func (c *tokenClaims) Valid() error {
	return c.StandardClaims.Valid()
}
