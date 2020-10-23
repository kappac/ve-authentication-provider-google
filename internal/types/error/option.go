package error

// OptionUpdater updates Error properties.
type OptionUpdater func(*veError)

// WithCode updates Code field of veError
func WithCode(c int32) OptionUpdater {
	return func(e *veError) {
		e.Code = c
	}
}

// WithDescription updates Description field of veError
func WithDescription(d string) OptionUpdater {
	return func(e *veError) {
		e.Description = d
	}
}
