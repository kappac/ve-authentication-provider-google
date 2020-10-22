package providerinfo

// VEProviderInfo contains an information about user,
// necessary to provide a service.
type VEProviderInfo interface {
	FullName() string
	GivenName() string
	FamilyName() string
	Picture() string
	Email() string
}

type veProviderInfo struct {
	PFullName   string `json:"full_name,omitempty"`
	PGivenName  string `json:"given_name,omitempty"`
	PFamilyName string `json:"family_name,omitempty"`
	PPicture    string `json:"picture,omitempty"`
	PEmail      string `json:"email,omitempty"`
}

// New creates new instance of VEProviderInfo.
func New(ous ...OptionUpdater) VEProviderInfo {
	pct := &veProviderInfo{}

	for _, ou := range ous {
		ou(pct)
	}

	return pct
}

func (pi *veProviderInfo) FullName() string {
	return pi.PFullName
}

func (pi *veProviderInfo) GivenName() string {
	return pi.PGivenName
}

func (pi *veProviderInfo) FamilyName() string {
	return pi.PFamilyName
}

func (pi *veProviderInfo) Picture() string {
	return pi.PPicture
}

func (pi *veProviderInfo) Email() string {
	return pi.PEmail
}
