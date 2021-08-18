package vote

type VoteInstruction uint16

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
