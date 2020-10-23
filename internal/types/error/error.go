package error

import (
	"fmt"

	"github.com/kappac/ve-authentication-provider-google/internal/pb"
	"github.com/kappac/ve-authentication-provider-google/internal/types"
	"github.com/kappac/ve-authentication-provider-google/internal/types/marshaller"
)

const (
	basicErrorCode = types.ConstErrorCodeTypesBasic + 100
	_              = iota + basicErrorCode
	errorCodeUnmarshalWrongType
)

var (
	errorUnmarshalWrongType = New(
		WithCode(errorCodeUnmarshalWrongType),
		WithDescription("A package provided for Unmarshal is of a wrong type"),
	)
)

// VEError is a basic error for VE project
type VEError interface {
	error
	marshaller.Marshaller

	GetCode() int32
	GetDescription() string
}

type veError struct {
	Code        int32  `json:"code"`
	Description string `json:"description"`
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
	return fmt.Sprintf("[%d]: %s", e.Code, e.Description)
}

func (e *veError) GetCode() int32 {
	return e.Code
}

func (e *veError) GetDescription() string {
	return e.Description
}

func (e *veError) Marshal() (interface{}, error) {
	p := &pb.VEError{
		Code:        e.GetCode(),
		Description: e.GetDescription(),
	}

	return p, nil
}

func (e *veError) Unmarshal(p interface{}) error {
	pbError, ok := p.(*pb.VEError)
	if !ok {
		return errorUnmarshalWrongType
	}

	e.Code = pbError.GetCode()
	e.Description = pbError.GetDescription()

	return nil
}
