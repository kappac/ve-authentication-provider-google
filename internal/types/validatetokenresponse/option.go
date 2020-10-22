package validatetokenresponse

import (
	"github.com/kappac/ve-authentication-provider-google/internal/types/error"
	"github.com/kappac/ve-authentication-provider-google/internal/types/providerinfo"
	"github.com/kappac/ve-authentication-provider-google/internal/types/validatetokenrequest"
)

// OptionUpdater updates veValidateTokenResponse properties.
type OptionUpdater func(*veValidateTokenResponse)

// WithInfo updates Info field of veValidateTokenResponse
func WithInfo(i providerinfo.VEProviderInfo) OptionUpdater {
	return func(tr *veValidateTokenResponse) {
		tr.PInfo = i
	}
}

// WithRequest updates Request field of veValidateTokenResponse
func WithRequest(r validatetokenrequest.VEValidateTokenRequest) OptionUpdater {
	return func(tr *veValidateTokenResponse) {
		tr.PRequest = r
	}
}

// WithError updates Error field of veValidateTokenResponse
func WithError(e error.VEError) OptionUpdater {
	return func(tr *veValidateTokenResponse) {
		tr.PError = e
	}
}
