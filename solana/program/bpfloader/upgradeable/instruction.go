package upgradable_loader

type InstructionID uint32

const (
	// Initialize a Buffer account
	InitializeBuffer InstructionID = iota

	// Write program data into a Buffer account
	Write

	// Deploy an executable program
	DeployWithMaxDataLen

	// Upgrade a program
	Upgrade

	// Set a new authority that is allowed to write the buffer or upgrade the program
	SetAuthority

	// Closes an account owned by the upgradeable loader of all lamports and withdraws all the lamports
	Close
)

type WriteInstruction struct {
	Offset uint32
	Bytes  []byte
}

type DeployWithMaxDataLenInstruction struct {
	MaxDataLen uint32
}
