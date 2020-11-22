package request

import (
	"github.com/kappac/ve-authentication-provider-google/internal/constants"
	veerror "github.com/kappac/ve-back-end-utils/pkg/error"
)

const (
	basicErrorCode = constants.ServiceBasicErrorCode + 300
	_              = iota + basicErrorCode
	errorCodeUnmarshalWrongType
	errorCodeVerifyTokenAbsent
)

var (
	errorUnmarshalWrongType = veerror.New(
		veerror.WithCode(errorCodeUnmarshalWrongType),
		veerror.WithMessage("A package provided for Unmarshal is of a wrong type"),
	)
	errorVerifyTokenAbsent = veerror.New(
		veerror.WithCode(errorCodeVerifyTokenAbsent),
		veerror.WithMessage("Token is absent"),
	)
)
