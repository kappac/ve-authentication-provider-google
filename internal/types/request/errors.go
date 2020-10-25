package request

import (
	"github.com/kappac/ve-authentication-provider-google/internal/types/constants"
	veerror "github.com/kappac/ve-authentication-provider-google/internal/types/error"
)

const (
	basicErrorCode = constants.TypesBasicErrorCode + 300
	_              = iota + basicErrorCode
	errorCodeUnmarshalWrongType
	errorCodeVerifyTokenAbsent
)

var (
	errorUnmarshalWrongType = veerror.New(
		veerror.WithCode(errorCodeUnmarshalWrongType),
		veerror.WithDescription("A package provided for Unmarshal is of a wrong type"),
	)
	errorVerifyTokenAbsent = veerror.New(
		veerror.WithCode(errorCodeVerifyTokenAbsent),
		veerror.WithDescription("Token is absent"),
	)
)
