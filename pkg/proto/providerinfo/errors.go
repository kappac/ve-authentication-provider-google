package providerinfo

import (
	"github.com/kappac/ve-authentication-provider-google/internal/constants"
	veerror "github.com/kappac/ve-back-end-utils/pkg/error"
)

const (
	basicErrorCode = constants.ServiceBasicErrorCode + 200
	_              = iota + basicErrorCode
	errorCodeUnmarshalWrongType
)

var (
	errorUnmarshalWrongType = veerror.New(
		veerror.WithCode(errorCodeUnmarshalWrongType),
		veerror.WithMessage("A package provided for Unmarshal is of a wrong type"),
	)
)
