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
	accounts []string,
	instruction InitializeAccountInstruction,
) ParsedInitializeAccount {
	return ParsedInitializeAccount{
		VoteAccount:          accounts[0],
		RentSysvar:           accounts[1],
		ClockSysvar:          accounts[2],
		Node:                 accounts[3],
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
	accounts []string,
	instruction AuthorizeInstruction,
) ParsedAuthorize {
	return ParsedAuthorize{
		VoteAccount:   accounts[0],
		ClockSysvar:   accounts[1],
		Authority:     accounts[2],
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
	accounts []string,
	instruction VoteInstruction,
) ParsedVote {
	return ParsedVote{
		VoteAccount:      accounts[0],
		SlotHashesSysvar: accounts[1],
		ClockSysvar:      accounts[2],
		VoteAuthority:    accounts[3],
		Vote:             NewParsedVoteData(instruction.Vote),
	}
}

//____________________________________________________________________________

type ParsedWithdraw struct {
	VoteAccount string `json:"voteAccount"`
	Destination string `json:"destination"`
	Lamports    uint64 `json:"lamports"`
}

func NewParsedWithdraw(
	accounts []string,
	instruction WithdrawInstruction,
) ParsedWithdraw {
	return ParsedWithdraw{
		VoteAccount: accounts[0],
		Destination: accounts[1],
		Lamports:    instruction.Amount,
	}
}

//____________________________________________________________________________

type ParsedUpdateValidatorIdentity struct {
	VoteAccount          string `json:"voteAccount"`
	NewValidatorIdentity string `json:"newValidatorIdentity"`
	WithdrawAuthority    string `json:"withdrawAuthority"`
}

func NewParsedUpdateValidatorIdentity(
	accounts []string,
) ParsedUpdateValidatorIdentity {
	return ParsedUpdateValidatorIdentity{
		VoteAccount:          accounts[0],
		NewValidatorIdentity: accounts[1],
		WithdrawAuthority:    accounts[2],
	}
}

//____________________________________________________________________________

type ParsedUpdateCommission struct {
	VoteAccount       string `json:"voteAccount"`
	WithdrawAuthority string `json:"withdrawAuthority"`
	Commission        uint8  `json:"commission"`
}

func NewParsedUpdateCommission(
	accounts []string,
	instruction UpdateCommissionInstruction,
) ParsedUpdateCommission {
	return ParsedUpdateCommission{
		VoteAccount:       accounts[0],
		WithdrawAuthority: accounts[1],
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
	accounts []string,
	instruction VoteSwitchInstruction,
) ParsedVoteSwitch {
	return ParsedVoteSwitch{
		VoteAccount:      accounts[0],
		SlotHashesSysvar: accounts[1],
		ClockSysvar:      accounts[2],
		VoteAuthority:    accounts[3],
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
	accounts []string,
	instruction AuthorizeCheckedInstruction,
) ParsedAuthorizeChecked {
	return ParsedAuthorizeChecked{
		VoteAccount:   accounts[0],
		ClockSysvar:   accounts[1],
		Authority:     accounts[2],
		NewAuthority:  accounts[3],
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
