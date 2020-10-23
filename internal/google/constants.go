package google

import (
	"time"

	"github.com/kappac/ve-authentication-provider-google/internal/constants"
)

const (
	fetcherGetTimeout            = 10 * time.Second
	fetcherCertsURL              = "https://www.googleapis.com/oauth2/v1/certs"
	fetcherCacheControlHeaderKey = "Cache-Control"
	fetcherMaxAgeProperty        = "max-age"
)

const (
	packageBasicErrorCode = constants.ServiceBasicErrorCode + 2000
)
