package proto

// Marshaller used for governing the process of serializing
// to/from PB packets
type Marshaller interface {
	Marshal() (interface{}, error)
	Unmarshal(interface{}) error
	Verify() error
}
