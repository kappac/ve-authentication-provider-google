package error

const (
	basicErrorCode = 100
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
