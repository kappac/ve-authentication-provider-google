package providerinfo

import (
	"github.com/kappac/ve-authentication-provider-google/internal/pb"
	"github.com/kappac/ve-authentication-provider-google/internal/types/marshaller"
)

// VEProviderInfo contains an information about user,
// necessary to provide a service.
type VEProviderInfo interface {
	marshaller.Marshaller

	GetFullName() string
	GetGivenName() string
	GetFamilyName() string
	GetPicture() string
	GetEmail() string
}

type veProviderInfo struct {
	FullName   string `json:"full_name,omitempty"`
	GivenName  string `json:"given_name,omitempty"`
	FamilyName string `json:"family_name,omitempty"`
	Picture    string `json:"picture,omitempty"`
	Email      string `json:"email,omitempty"`
}

// New creates new instance of VEProviderInfo.
func New(ous ...OptionUpdater) VEProviderInfo {
	pi := &veProviderInfo{}

	for _, ou := range ous {
		ou(pi)
	}

	return pi
}

func (pi *veProviderInfo) GetFullName() string {
	return pi.FullName
}

func (pi *veProviderInfo) GetGivenName() string {
	return pi.GivenName
}

func (pi *veProviderInfo) GetFamilyName() string {
	return pi.FamilyName
}

func (pi *veProviderInfo) GetPicture() string {
	return pi.Picture
}

func (pi *veProviderInfo) GetEmail() string {
	return pi.Email
}

func (pi *veProviderInfo) Marshal() (interface{}, error) {
	p := &pb.VEProviderInfo{
		FullName:   pi.GetFullName(),
		GivenName:  pi.GetGivenName(),
		FamilyName: pi.GetFamilyName(),
		Picture:    pi.GetPicture(),
		Email:      pi.GetEmail(),
	}

	return p, nil
}

func (pi *veProviderInfo) Unmarshal(p interface{}) error {
	pbInfo, ok := p.(*pb.VEProviderInfo)
	if !ok {
		return errorUnmarshalWrongType
	}

	pi.FullName = pbInfo.GetFullName()
	pi.GivenName = pbInfo.GetGivenName()
	pi.FamilyName = pbInfo.GetFamilyName()
	pi.Picture = pbInfo.GetPicture()
	pi.Email = pbInfo.GetEmail()

	return nil
}

func (pi *veProviderInfo) Verify() error {
	return nil
}
