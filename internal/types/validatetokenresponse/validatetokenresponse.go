package validatetokenresponse

import (
	"github.com/kappac/ve-authentication-provider-google/internal/pb"
	"github.com/kappac/ve-authentication-provider-google/internal/types/constants"
	veerror "github.com/kappac/ve-authentication-provider-google/internal/types/error"
	"github.com/kappac/ve-authentication-provider-google/internal/types/marshaller"
	"github.com/kappac/ve-authentication-provider-google/internal/types/providerinfo"
	"github.com/kappac/ve-authentication-provider-google/internal/types/validatetokenrequest"
)

const (
	basicErrorCode       = constants.TypesBasicErrorCode + 400
	errorCodeMarshalling = iota + basicErrorCode
	errorCodeUnmarshalWrongType
)

var (
	errorMarshaling = veerror.New(
		veerror.WithCode(errorCodeMarshalling),
		veerror.WithDescription("An error during parsing fields"),
	)
	errorUnmarshalWrongType = veerror.New(
		veerror.WithCode(errorCodeUnmarshalWrongType),
		veerror.WithDescription("A package provided for Unmarshal is of a wrong type"),
	)
)

// VEValidateTokenResponse is a wrapper for proto response.
type VEValidateTokenResponse interface {
	marshaller.Marshaller

	GetInfo() providerinfo.VEProviderInfo
	GetRequest() validatetokenrequest.VEValidateTokenRequest
	GetError() veerror.VEError
}

type veValidateTokenResponse struct {
	Info    providerinfo.VEProviderInfo                 `json:"info,omitempty"`
	Request validatetokenrequest.VEValidateTokenRequest `json:"request"`
	Error   veerror.VEError                             `json:"error,omitempty"`
}

func (tr *veValidateTokenResponse) GetInfo() providerinfo.VEProviderInfo {
	return tr.Info
}

func (tr *veValidateTokenResponse) GetRequest() validatetokenrequest.VEValidateTokenRequest {
	return tr.Request
}

func (tr *veValidateTokenResponse) GetError() veerror.VEError {
	return tr.Error
}

func (tr *veValidateTokenResponse) Marshal() (interface{}, error) {
	var (
		reqpb, infopb, errpb    interface{}
		reqErr, infoErr, errErr error
	)

	if req := tr.GetRequest(); req != nil {
		reqpb, reqErr = req.Marshal()
	}

	if info := tr.GetInfo(); info != nil {
		infopb, infoErr = info.Marshal()
	}

	if err := tr.GetError(); err != nil {
		errpb, errErr = err.Marshal()
	}

	if reqErr == nil || infoErr == nil || errErr == nil {
		return nil, errorMarshaling
	}

	p := &pb.VEValidateTokenResponse{
		Request: reqpb.(*pb.VEValidateTokenRequest),
		Info:    infopb.(*pb.VEProviderInfo),
		Error:   errpb.(*pb.VEError),
	}

	return p, nil
}
