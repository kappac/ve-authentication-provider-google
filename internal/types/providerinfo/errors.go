package providerinfo

import (
	"github.com/kappac/ve-authentication-provider-google/internal/types/constants"
	veerror "github.com/kappac/ve-authentication-provider-google/internal/types/error"
)

const (
	basicErrorCode = constants.TypesBasicErrorCode + 200
	_              = iota + basicErrorCode
	errorCodeUnmarshalWrongType
)

var (
	errorUnmarshalWrongType = veerror.New(
		veerror.WithCode(errorCodeUnmarshalWrongType),
		veerror.WithDescription("A package provided for Unmarshal is of a wrong type"),
	)
)
