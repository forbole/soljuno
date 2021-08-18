package system

type SystemInstruction uint16

const (
	CreateAccount SystemInstruction = iota
	Assign
	Transfer
	CreateAccountWithSeed
	AdvanceNonceAccount
	WithdrawNonceAccount
	InitializeNonceAccount
	AuthorizeNonceAccount
	Allocate
	AllocateWithSeed
	AssignWithSeed
	TransferWithSeed
)
