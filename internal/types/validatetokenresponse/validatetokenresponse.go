package validatetokenresponse

import (
	"github.com/kappac/ve-authentication-provider-google/internal/types/error"
	"github.com/kappac/ve-authentication-provider-google/internal/types/providerinfo"
	"github.com/kappac/ve-authentication-provider-google/internal/types/validatetokenrequest"
)

// VEValidateTokenResponse is a wrapper for proto response.
type VEValidateTokenResponse interface {
	Info() providerinfo.VEProviderInfo
	Request() validatetokenrequest.VEValidateTokenRequest
	Error() error.VEError
}

type veValidateTokenResponse struct {
	PInfo    providerinfo.VEProviderInfo                 `json:"info,omitempty"`
	PRequest validatetokenrequest.VEValidateTokenRequest `json:"request"`
	PError   error.VEError                               `json:"error,omitempty"`
}
