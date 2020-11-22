package google

import (
	veerror "github.com/kappac/ve-back-end-utils/pkg/error"
)

const (
	certificatesBasicErrorCode = packageBasicErrorCode + 200

	errorCodeNoCertificate = iota + certificatesBasicErrorCode
)

var (
	errorNoCertificate = veerror.New(
		veerror.WithCode(errorCodeNoCertificate),
		veerror.WithMessage("A certificate is absent"),
	)
)
