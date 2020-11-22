package request

import (
	"github.com/kappac/ve-authentication-provider-google/internal/pb"
	"github.com/kappac/ve-back-end-utils/pkg/proto"
)

// VEValidateTokenRequest is a wrapper for proto request.
type VEValidateTokenRequest interface {
	proto.Marshaller

	GetToken() string
}

type veValidateTokenRequest struct {
	Token string `json:"token"`
}

// New creates new instance of VEValidateTokenRequest.
func New(ous ...OptionUpdater) VEValidateTokenRequest {
	pct := &veValidateTokenRequest{}

	for _, ou := range ous {
		ou(pct)
	}

	return pct
}

func (tr *veValidateTokenRequest) GetToken() string {
	return tr.Token
}

func (tr *veValidateTokenRequest) Marshal() (interface{}, error) {
	p := &pb.VEValidateTokenRequest{
		Token: tr.GetToken(),
	}

	return p, nil
}

func (tr *veValidateTokenRequest) Unmarshal(p interface{}) error {
	pbRequest, ok := p.(*pb.VEValidateTokenRequest)
	if !ok {
		return errorUnmarshalWrongType
	}

	tr.Token = pbRequest.GetToken()

	return nil
}

func (tr *veValidateTokenRequest) Verify() error {
	if tr.GetToken() == "" {
		return errorVerifyTokenAbsent
	}

	return nil
}
