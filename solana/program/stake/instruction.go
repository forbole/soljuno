package stake

type StakeInstruction uint32

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
