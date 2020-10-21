package server

import "fmt"

type VEError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func (e VEError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

type VETokenInfo struct {
	FullName   string `json:"full_name,omitempty"`
	GivenName  string `json:"given_name,omitempty"`
	FamilyName string `json:"family_name,omitempty"`
	Picture    string `json:"picture,omitempty"`
	Email      string `json:"email,omitempty"`
}

type VEValidateTokenRequest struct {
	Token string `json:"token"`
}

type VEValidateTokenResponse struct {
	Info    VETokenInfo            `json:"info,omitempty"`
	Request VEValidateTokenRequest `json:"request"`
	Error   VEError                `json:"error,omitempty"`
}
