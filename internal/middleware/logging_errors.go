package middleware

import (
	veerror "github.com/kappac/ve-authentication-provider-google/internal/types/error"
)

const (
	loggingBasicErrorCode     = middlewareBasicErrorCode + 100
	errorCodeLoggingNoService = iota + loggingBasicErrorCode
)

var (
	errorLoggingNoService = veerror.New(
		veerror.WithCode(errorCodeLoggingNoService),
		veerror.WithDescription("No service instance provided for Logging"),
	)
)
