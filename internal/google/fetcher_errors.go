package google

import (
	"fmt"

	veerror "github.com/kappac/ve-authentication-provider-google/internal/types/error"
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
	errorCacheControlAbscent = veerror.New(
		veerror.WithCode(errorCodeCacheControlHeaderAbscent),
		veerror.WithDescription(
			fmt.Sprintf("\"%s\" is absent in the response header", fetcherCacheControlHeaderKey),
		),
	)
	errorMaxAgePropertyAbscent = veerror.New(
		veerror.WithCode(errorCodeMaxAgePropertyAbsent),
		veerror.WithDescription(
			fmt.Sprintf("\"%s\" is absent in the response header", fetcherMaxAgeProperty),
		),
	)
	errorMaxAgeValueAbscent = veerror.New(
		veerror.WithCode(errorCodeMaxAgeValueAbsent),
		veerror.WithDescription(
			fmt.Sprintf("Value for \"%s\" is absent", fetcherMaxAgeProperty),
		),
	)
	errorMaxAgeConvertion = veerror.New(
		veerror.WithCode(errorCodeMaxAgeConvertion),
		veerror.WithDescription(
			fmt.Sprintf("Error during \"%s\" convertion", fetcherMaxAgeProperty),
		),
	)
	errorJSONDecode = veerror.New(
		veerror.WithCode(errorCodeJSONDecode),
		veerror.WithDescription("Error during response decoding"),
	)
)
