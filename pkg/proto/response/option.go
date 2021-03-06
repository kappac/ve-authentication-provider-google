package response

import (
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/providerinfo"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/request"
	veerror "github.com/kappac/ve-back-end-utils/pkg/error"
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
func WithError(e veerror.Error) OptionUpdater {
	return func(tr *veValidateTokenResponse) {
		tr.Error = e
	}
}
