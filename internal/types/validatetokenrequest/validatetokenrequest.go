package validatetokenrequest

import (
	"github.com/kappac/ve-authentication-provider-google/internal/pb"
	"github.com/kappac/ve-authentication-provider-google/internal/types"
	veerror "github.com/kappac/ve-authentication-provider-google/internal/types/error"
	"github.com/kappac/ve-authentication-provider-google/internal/types/marshaller"
)

const (
	basicErrorCode = types.ConstErrorCodeTypesBasic + 300
	_              = iota + basicErrorCode
	errorCodeUnmarshalWrongType
)

var (
	errorUnmarshalWrongType = veerror.New(
		veerror.WithCode(errorCodeUnmarshalWrongType),
		veerror.WithDescription("A package provided for Unmarshal is of a wrong type"),
	)
)

// VEValidateTokenRequest is a wrapper for proto request.
type VEValidateTokenRequest interface {
	marshaller.Marshaller

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
