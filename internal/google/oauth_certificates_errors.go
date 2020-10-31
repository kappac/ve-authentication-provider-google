package google

import (
	veerror "github.com/kappac/ve-authentication-provider-google/pkg/proto/error"
)

const (
	certificatesBasicErrorCode = packageBasicErrorCode + 200

	errorCodeNoCertificate = iota + certificatesBasicErrorCode
)

var (
	errorNoCertificate = veerror.New(
		veerror.WithCode(errorCodeNoCertificate),
		veerror.WithDescription("A certificate is absent"),
	)
)
