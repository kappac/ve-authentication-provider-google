package response

import (
	"github.com/kappac/ve-authentication-provider-google/pkg/error"
	"github.com/kappac/ve-authentication-provider-google/pkg/providerinfo"
	"github.com/kappac/ve-authentication-provider-google/pkg/request"
)

// OptionUpdater updates veValidateTokenResponse properties.
type OptionUpdater func(*veValidateTokenResponse)

// WithInfo updates Info field of veValidateTokenResponse
func WithInfo(i providerinfo.VEProviderInfo) OptionUpdater {
	return func(tr *veValidateTokenResponse) {
		tr.Info = i
	}
}

// WithRequest updates Request field of veValidateTokenResponse
func WithRequest(r request.VEValidateTokenRequest) OptionUpdater {
	return func(tr *veValidateTokenResponse) {
		tr.Request = r
	}
}

// WithError updates Error field of veValidateTokenResponse
func WithError(e error.VEError) OptionUpdater {
	return func(tr *veValidateTokenResponse) {
		tr.Error = e
	}
}
