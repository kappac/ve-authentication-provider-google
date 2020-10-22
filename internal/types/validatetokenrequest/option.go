package validatetokenrequest

// OptionUpdater updates Error properties.
type OptionUpdater func(*veValidateTokenRequest)

// WithToken updates Token field of an veValidateTokenRequest
func WithToken(t string) OptionUpdater {
	return func(tr *veValidateTokenRequest) {
		tr.PToken = t
	}
}
