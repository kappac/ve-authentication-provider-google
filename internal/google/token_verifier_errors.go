package google

import (
	veerror "github.com/kappac/ve-authentication-provider-google/pkg/error"
)

const (
	tokenVerifierBasicErrorCode = packageBasicErrorCode + 300

	errorCodeNoCertificateKeyID = iota + tokenVerifierBasicErrorCode
)

var (
	errorNoCertificateKeyID = veerror.New(
		veerror.WithCode(errorCodeNoCertificateKeyID),
		veerror.WithDescription("Certificate key id has not been found"),
	)
)
