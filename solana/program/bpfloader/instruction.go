package bpfloader

type LoaderInstruction uint32

const (
	// Write program data into an account
	Write LoaderInstruction = iota
	Finalize
)

// Finalize an account loaded with program data for execution
type WriteInstruction struct {
	Offset uint32
	Bytes  []byte
}
