package statusservice

import (
	veerror "github.com/kappac/ve-authentication-provider-google/pkg/error"
)

const (
	basicErrorCode              = 200
	errorCodeNoProbesConfigured = iota + basicErrorCode
)

var (
	errorNoProbesConfigured = veerror.New(
		veerror.WithCode(errorCodeNoProbesConfigured),
		veerror.WithDescription("At least one probe source should configured"),
	)
)
