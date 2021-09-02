package vote

type ParsedInitializeAccount struct {
	VoteAccount          string `json:"voteAccount"`
	RentSysvar           string `json:"rentSysvar"`
	ClockSysvar          string `json:"clockSysvar"`
	Node                 string `json:"node"`
	AuthorizedVoter      string `json:"authorizedVoter"`
	AuthorizedWithdrawer string `json:"authorizedWithdrawer"`
	Commission           uint8  `json:"commission"`
}

func NewParsedInitializeAccount(
	voteAccount,
	rentSysvar,
	clockSysvar,
	node string,
	instruction InitializeAccountInstruction,
) ParsedInitializeAccount {
	return ParsedInitializeAccount{
		VoteAccount:          voteAccount,
		RentSysvar:           rentSysvar,
		ClockSysvar:          clockSysvar,
		Node:                 node,
		AuthorizedVoter:      instruction.VoteInit.AuthorizedVoter.String(),
		AuthorizedWithdrawer: instruction.VoteInit.AuthorizedWithdrawer.String(),
		Commission:           instruction.VoteInit.Commission,
	}
}

//____________________________________________________________________________

type ParsedAuthorize struct {
	VoteAccount   string `json:"voteAccount"`
	ClockSysvar   string `json:"clockSysvar"`
	Authority     string `json:"authority"`
	NewAuthority  string `json:"newAuthority"`
	AuthorityType string `json:"authorityType"`
}

func NewParsedAuthorize(
	voteAccount,
	clockSysvar,
	authority string,
	instruction AuthorizeInstruction,
) ParsedAuthorize {
	return ParsedAuthorize{
		VoteAccount:   voteAccount,
		ClockSysvar:   clockSysvar,
		Authority:     authority,
		NewAuthority:  instruction.Pubkey.String(),
		AuthorityType: NewParsedAuthorityType(instruction.VoteAuthorize),
	}
}

//____________________________________________________________________________

type ParsedVote struct {
	VoteAccount      string         `json:"voteAccount"`
	SlotHashesSysvar string         `json:"slotHashesSysvar"`
	ClockSysvar      string         `json:"clockSysvar"`
	VoteAuthority    string         `json:"voteAuthority"`
	Vote             ParsedVoteData `json:"vote"`
}

func NewParsedVote(
	voteAccount,
	slotHashesSysvar,
	clockSysvar,
	voteAuthority string,
	instruction VoteInstruction,
) ParsedVote {
	return ParsedVote{
		VoteAccount:      voteAccount,
		SlotHashesSysvar: slotHashesSysvar,
		ClockSysvar:      clockSysvar,
		VoteAuthority:    voteAuthority,
		Vote:             NewParsedVoteData(instruction.Vote),
	}
}

//____________________________________________________________________________

type ParsedWithdraw struct {
	VoteAccount       string `json:"voteAccount"`
	Destination       string `json:"destination"`
	WithdrawAuthority string `json:"withdrawAuthority"`
	Lamports          uint64 `json:"lamports"`
}

func NewParsedWithdraw(
	voteAccount,
	destination,
	withdrawAuthority string,
	instruction WithdrawInstruction,
) ParsedWithdraw {
	return ParsedWithdraw{
		VoteAccount:       voteAccount,
		Destination:       destination,
		WithdrawAuthority: withdrawAuthority,
		Lamports:          instruction.Amount,
	}
}

//____________________________________________________________________________

type ParsedUpdateValidatorIdentity struct {
	VoteAccount          string `json:"voteAccount"`
	NewValidatorIdentity string `json:"newValidatorIdentity"`
	WithdrawAuthority    string `json:"withdrawAuthority"`
}

func NewParsedUpdateValidatorIdentity(
	voteAccount,
	newValidatorIdentity,
	withdrawAuthority string,
) ParsedUpdateValidatorIdentity {
	return ParsedUpdateValidatorIdentity{
		VoteAccount:          voteAccount,
		NewValidatorIdentity: newValidatorIdentity,
		WithdrawAuthority:    withdrawAuthority,
	}
}

//____________________________________________________________________________

type ParsedUpdateCommission struct {
	VoteAccount       string `json:"voteAccount"`
	WithdrawAuthority string `json:"withdrawAuthority"`
	Commission        uint8  `json:"commission"`
}

func NewParsedUpdateCommission(
	voteAccount,
	withdrawAuthority string,
	instruction UpdateCommissionInstruction,
) ParsedUpdateCommission {
	return ParsedUpdateCommission{
		VoteAccount:       voteAccount,
		WithdrawAuthority: withdrawAuthority,
		Commission:        instruction.Commission,
	}
}

//____________________________________________________________________________

type ParsedVoteSwitch struct {
	VoteAccount      string         `json:"voteAccount"`
	SlotHashesSysvar string         `json:"slotHashesSysvar"`
	ClockSysvar      string         `json:"clockSysvar"`
	VoteAuthority    string         `json:"voteAuthority"`
	Vote             ParsedVoteData `json:"vote"`
	Hash             string         `json:"hash"`
}

func NewParsedVoteSwitch(
	voteAccount,
	slotHashesSysvar,
	clockSysvar,
	voteAuthority string,
	instruction VoteSwitchInstruction,
) ParsedVoteSwitch {
	return ParsedVoteSwitch{
		VoteAccount:      voteAccount,
		SlotHashesSysvar: slotHashesSysvar,
		ClockSysvar:      clockSysvar,
		VoteAuthority:    voteAuthority,
		Vote:             NewParsedVoteData(instruction.Vote),
		Hash:             instruction.Hash.String(),
	}
}

//____________________________________________________________________________

type ParsedAuthorizeChecked struct {
	VoteAccount   string `json:"voteAccount"`
	ClockSysvar   string `json:"clockSysvar"`
	Authority     string `json:"authority"`
	NewAuthority  string `json:"newAuthority"`
	AuthorityType string `json:"authorityType"`
}

func NewParsedAuthorizeChecked(
	voteAccount,
	clockSysvar,
	authority,
	newAuthority string,
	instruction AuthorizeCheckedInstruction,
) ParsedAuthorizeChecked {
	return ParsedAuthorizeChecked{
		VoteAccount:   voteAccount,
		ClockSysvar:   clockSysvar,
		Authority:     authority,
		NewAuthority:  newAuthority,
		AuthorityType: NewParsedAuthorityType(instruction.VoteAuthorize),
	}
}

// ____________________________________________________________________________

// The instances used in parsed instructions

type ParsedVoteData struct {
	Slots     []uint64 `json:"slots"`
	Hash      string   `json:"hash"`
	Timestamp *uint64  `json:"timestamp,omitempty"`
}

func NewParsedVoteData(voteData VoteData) ParsedVoteData {
	return ParsedVoteData{
		Slots:     voteData.Slots,
		Hash:      voteData.Hash.String(),
		Timestamp: voteData.Timestamp,
	}
}

func NewParsedAuthorityType(typ VoteAuthorize) string {
	switch typ {
	case Voter:
		return "voter"
	case Withdrawer:
		return "withdrawer"
	}
	return ""
}
