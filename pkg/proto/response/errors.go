package response

import (
	"github.com/kappac/ve-authentication-provider-google/internal/constants"
	veerror "github.com/kappac/ve-back-end-utils/pkg/proto/error"
)

const (
	basicErrorCode       = constants.ServiceBasicErrorCode + 400
	errorCodeMarshalling = iota + basicErrorCode
	errorCodeUnmarshalWrongType
	errorCodeVerifyRequestAbsent
	errorCodeVerifyInfoErrorAbsent
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
	errorVerifyRequestAbsent = veerror.New(
		veerror.WithCode(errorCodeVerifyRequestAbsent),
		veerror.WithDescription("Request is absent"),
	)
	errorVerifyInfoErrorAbsent = veerror.New(
		veerror.WithCode(errorCodeVerifyInfoErrorAbsent),
		veerror.WithDescription("Info or Error is absent"),
	)
)
