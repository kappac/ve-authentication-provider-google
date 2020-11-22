package google

import (
	"fmt"

	veerror "github.com/kappac/ve-back-end-utils/pkg/error"
)

const (
	fetcherBasicErrorCode = packageBasicErrorCode + 100

	errorCodeCacheControlHeaderAbsent = iota + fetcherBasicErrorCode
	errorCodeMaxAgePropertyAbsent
	errorCodeMaxAgeValueAbsent
	errorCodeMaxAgeConvertion
	errorCodeJSONDecode
)

var (
	errorCacheControlAbsent = veerror.New(
		veerror.WithCode(errorCodeCacheControlHeaderAbsent),
		veerror.WithMessage(
			fmt.Sprintf("\"%s\" is absent in the response header", fetcherCacheControlHeaderKey),
		),
	)
	errorMaxAgePropertyAbsent = veerror.New(
		veerror.WithCode(errorCodeMaxAgePropertyAbsent),
		veerror.WithMessage(
			fmt.Sprintf("\"%s\" is absent in the response header", fetcherMaxAgeProperty),
		),
	)
	errorMaxAgeValueAbsent = veerror.New(
		veerror.WithCode(errorCodeMaxAgeValueAbsent),
		veerror.WithMessage(
			fmt.Sprintf("Value for \"%s\" is absent", fetcherMaxAgeProperty),
		),
	)
	errorMaxAgeConvertion = veerror.New(
		veerror.WithCode(errorCodeMaxAgeConvertion),
		veerror.WithMessage(
			fmt.Sprintf("Error during \"%s\" convertion", fetcherMaxAgeProperty),
		),
	)
	errorJSONDecode = veerror.New(
		veerror.WithCode(errorCodeJSONDecode),
		veerror.WithMessage("Error during response decoding"),
	)
)
