package stake

type StakeInstruction uint16

const (
	Initialize StakeInstruction = iota
	Authorize
	DelegateStake
	Split
	Withdraw
	Deactivate
	SetLockup
	Merge
	AuthorizeWithSeed
	InitializeChecked
	AuthorizeChecked
	AuthorizeCheckedWithSeed
	SetLockupChecked
)
