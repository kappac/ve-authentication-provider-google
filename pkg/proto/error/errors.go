package error

import (
	"github.com/kappac/ve-authentication-provider-google/pkg/proto"
)

const (
	basicErrorCode = proto.BasicErrorCode + 1000
	_              = iota + basicErrorCode
	errorCodeUnmarshalWrongType
	errorCodeVerifyDescribtionAbsent
)

var (
	errorUnmarshalWrongType = New(
		WithCode(errorCodeUnmarshalWrongType),
		WithDescription("A package provided for Unmarshal is of a wrong type"),
	)
	errorVerifyDescribtionAbsent = New(
		WithCode(errorCodeVerifyDescribtionAbsent),
		WithDescription("Description is absent"),
	)
)
