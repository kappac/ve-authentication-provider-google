package validatetokenrequest

// VEValidateTokenRequest is a wrapper for proto request.
type VEValidateTokenRequest interface {
	Token() string
}

type veValidateTokenRequest struct {
	PToken string `json:"token"`
}

// New creates new instance of VEValidateTokenRequest.
func New(ous ...OptionUpdater) VEValidateTokenRequest {
	pct := &veValidateTokenRequest{}

	for _, ou := range ous {
		ou(pct)
	}

	return pct
}

func (tr *veValidateTokenRequest) Token() string {
	return tr.PToken
}
