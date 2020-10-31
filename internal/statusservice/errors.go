package statusservice

import (
	veerror "github.com/kappac/ve-authentication-provider-google/pkg/proto/error"
)

const (
	basicErrorCode = 200

	errorCodeNoProbesConfigured = iota + basicErrorCode
	errorCodeNoSourceData
)

var (
	errorNoProbesConfigured = veerror.New(
		veerror.WithCode(errorCodeNoProbesConfigured),
		veerror.WithDescription("At least one probe source should configured"),
	)
	errorNoSourceData = veerror.New(
		veerror.WithCode(errorCodeNoSourceData),
		veerror.WithDescription("No source data for a probe type"),
	)
)
