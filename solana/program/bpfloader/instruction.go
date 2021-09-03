package bpfloader

type InstructionID uint32

const (
	// Write program data into an account
	Write InstructionID = iota

	// Finalize an account loaded with program data for execution
	Finalize
)

type WriteInstruction struct {
	Offset uint32
	Bytes  []byte
}
