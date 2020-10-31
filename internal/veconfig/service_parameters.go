package veconfig

// ServiceParameters describes what methods should has
// a type to be a ServiceParameters object,
type ServiceParameters interface {
	Verify() error
}

// DefaultServiceParameters represents parameters that are common for
// all the VE services.
type DefaultServiceParameters struct {
	Name          string `json:"NAME,omitempty"`
	Address       string `json:"ADDRESS,omitempty"`
	ProbesAddress string `json:"PROBES_ADDRESS,omitempty"`
}

// Verify checks if all the default params are valid.
func (p DefaultServiceParameters) Verify() error {
	return nil
}
