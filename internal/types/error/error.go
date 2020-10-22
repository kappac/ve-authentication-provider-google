package error

import (
	"fmt"
)

// VEError is a basic error for VE project
type VEError interface {
	error

	Code() int32
	Description() string
}

type veError struct {
	PCode        int32  `json:"code"`
	PDescription string `json:"description"`
}

// New creates new instance of VEError.
func New(ous ...OptionUpdater) VEError {
	pct := &veError{}

	for _, ou := range ous {
		ou(pct)
	}

	return pct
}

func (e *veError) Error() string {
	return fmt.Sprintf("[%d]: %s", e.PCode, e.PDescription)
}

func (e *veError) Code() int32 {
	return e.PCode
}

func (e *veError) Description() string {
	return e.PDescription
}
