package google

import (
	veerror "github.com/kappac/ve-back-end-utils/pkg/error"
)

const (
	tokenVerifierBasicErrorCode = packageBasicErrorCode + 300

	errorCodeNoCertificateKeyID = iota + tokenVerifierBasicErrorCode
)

var (
	errorNoCertificateKeyID = veerror.New(
		veerror.WithCode(errorCodeNoCertificateKeyID),
		veerror.WithMessage("Certificate key id has not been found"),
	)
)
