package veconfig

const (
	globalEnvPrefix = "VE_"
)

// VEGlobalParameters ...
type VEGlobalParameters struct {
	Debug bool `json:"DEBUG"`
}

// Verify checks if VEGlobalParameters has correct values.
func (gp VEGlobalParameters) Verify() error {
	return nil
}

var (
	// Global is a set of VE global parameters
	Global VEGlobalParameters = VEGlobalParameters{
		Debug: true,
	}
)

func init() {
	ParsePrefix(globalEnvPrefix, &Global)
}
