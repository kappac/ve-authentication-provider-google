package response

import (
	"github.com/kappac/ve-authentication-provider-google/internal/pb"
	"github.com/kappac/ve-authentication-provider-google/internal/types/constants"
	veerror "github.com/kappac/ve-authentication-provider-google/internal/types/error"
	"github.com/kappac/ve-authentication-provider-google/internal/types/marshaller"
	"github.com/kappac/ve-authentication-provider-google/internal/types/providerinfo"
	"github.com/kappac/ve-authentication-provider-google/internal/types/request"
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
	GetRequest() request.VEValidateTokenRequest
	GetError() veerror.VEError
}

type veValidateTokenResponse struct {
	Info    providerinfo.VEProviderInfo    `json:"info,omitempty"`
	Request request.VEValidateTokenRequest `json:"request"`
	Error   veerror.VEError                `json:"error,omitempty"`
}

// New constructs VEValidateTokenResponse instance
func New(ous ...OptionUpdater) VEValidateTokenResponse {
	r := &veValidateTokenResponse{}

	for _, ou := range ous {
		ou(r)
	}

	return r
}

func (tr *veValidateTokenResponse) GetInfo() providerinfo.VEProviderInfo {
	return tr.Info
}

func (tr *veValidateTokenResponse) GetRequest() request.VEValidateTokenRequest {
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

func (tr *veValidateTokenResponse) Unmarshal(p interface{}) error {
	var (
		veinfo = providerinfo.New()
		vereq  = request.New()
		veerr  = veerror.New()
	)

	pbResponse, ok := p.(*pb.VEValidateTokenResponse)
	if !ok {
		return errorUnmarshalWrongType
	}

	if err := veinfo.Unmarshal(pbResponse.GetInfo()); err == nil {
		tr.Info = veinfo
	} else {
		return err
	}

	if err := vereq.Unmarshal(pbResponse.GetRequest()); err == nil {
		tr.Request = vereq
	} else {
		return err
	}

	if err := veerr.Unmarshal(pbResponse.GetError()); err == nil {
		tr.Error = veerr
	} else {
		return err
	}

	return nil
}
