package token

type TokenInstruction uint16

const (
	InitializeMint TokenInstruction = iota
	InitializeAccount
	InitializeMultisig
	Transfer
	Approve
	Revoke
	SetAuthority
	MintTo
	Burn
	CloseAccount
	FreezeAccount
	ThawAccount
	TransferChecked
	ApproveChecked
	MintToChecked
	BurnChecked
	InitializeAccount2
	SyncNative
)
