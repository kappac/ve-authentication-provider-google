package google

import (
	veerror "github.com/kappac/ve-back-end-utils/pkg/proto/error"
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
