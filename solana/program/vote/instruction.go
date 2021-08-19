package vote

type VoteInstruction uint32

const (
	InitializeAccount VoteInstruction = iota
	Authorize
	Vote
	Withdraw
	UpdateValidatorIdentity
	UpdateCommission
	VoteSwitch
	AuthorizeChecked
)
