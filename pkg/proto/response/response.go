package response

import (
	"github.com/kappac/ve-authentication-provider-google/internal/pb"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/providerinfo"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/request"
	"github.com/kappac/ve-back-end-utils/pkg/proto"
	veerror "github.com/kappac/ve-back-end-utils/pkg/proto/error"
)

// VEValidateTokenResponse is a wrapper for proto response.
type VEValidateTokenResponse interface {
	proto.Marshaller

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
		reqpb                   *pb.VEValidateTokenRequest
		infopb                  *pb.VEProviderInfo
		errpb                   *pb.VEError
		reqErr, infoErr, errErr error
	)

	if req := tr.GetRequest(); req != nil {
		req, reqErr := req.Marshal()

		if reqErr == nil {
			reqpb = req.(*pb.VEValidateTokenRequest)
		}
	}

	if info := tr.GetInfo(); info != nil {
		info, infoErr := info.Marshal()

		if infoErr == nil {
			infopb = info.(*pb.VEProviderInfo)
		}
	}

	if err := tr.GetError(); err != nil {
		err, errErr := err.Marshal()

		if errErr == nil {
			errpb = err.(*pb.VEError)
		}
	}

	if reqErr != nil && infoErr != nil && errErr != nil {
		return nil, errorMarshaling
	}

	p := &pb.VEValidateTokenResponse{
		Request: reqpb,
		Info:    infopb,
		Error:   errpb,
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

func (tr *veValidateTokenResponse) Verify() error {
	if tr.GetRequest() == nil {
		return errorVerifyRequestAbsent
	}

	if tr.GetInfo() == nil && tr.GetError() == nil {
		return errorVerifyInfoErrorAbsent
	}

	return nil
}
