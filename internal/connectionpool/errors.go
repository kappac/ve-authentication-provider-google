package connectionpool

import (
	veerror "github.com/kappac/ve-authentication-provider-google/pkg/proto/error"
)

const (
	basicErrorCode = 300

	errorCodeNoConstructor = iota + basicErrorCode
)

var (
	errorNoNoConstructor = veerror.New(
		veerror.WithCode(errorCodeNoConstructor),
		veerror.WithDescription("A connection constructor is not set"),
	)
)
