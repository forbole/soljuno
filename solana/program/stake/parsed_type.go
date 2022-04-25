package stake

type ParsedInitialize struct {
	StakeAccount string           `json:"stakeAccount"`
	RentSysvar   string           `json:"rentSysvar"`
	Authorized   ParsedAuthorized `json:"authorized"`
	Lockup       ParsedLockup     `json:"lockup"`
}

func NewParsedInitialize(
	accounts []string,
	instruction InitializeInstruction,
) ParsedInitialize {
	return ParsedInitialize{
		StakeAccount: accounts[0],
		RentSysvar:   accounts[1],
		Authorized:   NewParsedAuthorized(instruction.Authorized),
		Lockup:       NewParsedLockup(instruction.Lockup),
	}
}

//____________________________________________________________________________

type ParsedAuthorize struct {
	StakeAccount  string `json:"stakeAccount"`
	ClockSysvar   string `json:"clockSysvar"`
	Authority     string `json:"authority"`
	NewAuthority  string `json:"newAuthority"`
	AuthorityType string `json:"authorityType"`

	Custodian string `json:"custodian,omitempty"`
}

func NewParsedAuthorize(
	accounts []string,
	instruction AuthorizeInstruction,
) ParsedAuthorize {
	authorized := ParsedAuthorize{
		StakeAccount:  accounts[0],
		ClockSysvar:   accounts[1],
		Authority:     accounts[2],
		NewAuthority:  instruction.Pubkey.String(),
		AuthorityType: NewParsedStakeAuthorize(instruction.StakeAuthorize),
	}

	if len(accounts) >= 4 {
		authorized.Custodian = accounts[3]
	}

	return authorized
}

//____________________________________________________________________________

type ParsedDelegateStake struct {
	StakeAccount       string `json:"stakeAccount"`
	VoteAccount        string `json:"voteAccount"`
	ClockSysvar        string `json:"clockSysvar"`
	StakeHistorySysvar string `json:"stakeHistorySysvar"`
	StakeConfigAccount string `json:"stakeConfigAccount"`
	StakeAuthority     string `json:"stakeAuthority"`
}

func NewParsedDelegateStake(
	accounts []string,
) ParsedDelegateStake {
	return ParsedDelegateStake{
		StakeAccount:       accounts[0],
		VoteAccount:        accounts[1],
		ClockSysvar:        accounts[2],
		StakeHistorySysvar: accounts[3],
		StakeConfigAccount: accounts[4],
		StakeAuthority:     accounts[5],
	}
}

//____________________________________________________________________________

type ParsedSplit struct {
	StakeAccount    string `json:"stakeAccount"`
	NewSplitAccount string `json:"newSplitAccount"`
	StakeAuthority  string `json:"stakeAuthority"`
	Lamports        uint64 `json:"lamports"`
}

func NewParsedSplit(
	accounts []string,
	instruction SplitInstruction,
) ParsedSplit {
	return ParsedSplit{
		StakeAccount:    accounts[0],
		NewSplitAccount: accounts[1],
		StakeAuthority:  accounts[2],
		Lamports:        instruction.Lamports,
	}
}

//____________________________________________________________________________

type ParsedWithdraw struct {
	StakeAccount       string `json:"stakeAccount"`
	Destination        string `json:"destination"`
	ClockSysvar        string `json:"clockSysvar"`
	StakeHistorySysvar string `json:"stakeHistorySysvar"`
	WithdrawAuthority  string `json:"withdrawAuthority"`
	Lamports           uint64 `json:"lamports"`
}

func NewParsedWithdraw(
	accounts []string,
	instruction WithdrawInstruction,
) ParsedWithdraw {
	return ParsedWithdraw{
		StakeAccount:       accounts[0],
		Destination:        accounts[1],
		ClockSysvar:        accounts[2],
		StakeHistorySysvar: accounts[3],
		WithdrawAuthority:  accounts[4],
		Lamports:           instruction.Lamports,
	}
}

//____________________________________________________________________________

type ParsedDeactivate struct {
	StakeAccount   string `json:"stakeAccount"`
	ClockSysvar    string `json:"clockSysvar"`
	StakeAuthority string `json:"stakeAuthority"`
}

func NewParsedDeactivate(
	accounts []string,
) ParsedDeactivate {
	return ParsedDeactivate{
		StakeAccount:   accounts[0],
		ClockSysvar:    accounts[1],
		StakeAuthority: accounts[2],
	}
}

//____________________________________________________________________________

type ParsedSetLockup struct {
	StakeAccount string       `json:"stakeAccount"`
	Custodian    string       `json:"custodian"`
	Lockup       ParsedLockup `json:"lockup"`
}

func NewParsedSetLockup(
	accounts []string,
	instruction SetLockupInstruction,
) ParsedSetLockup {
	return ParsedSetLockup{
		StakeAccount: accounts[0],
		Custodian:    accounts[1],
		Lockup:       NewParsedLockupFromArgs(instruction.LockupArgs),
	}
}

//____________________________________________________________________________

type ParsedMerge struct {
	Destination        string `json:"destination"`
	Source             string `json:"source"`
	ClockSysvar        string `json:"clockSysvar"`
	StakeHistorySysvar string `json:"stakeHistorySysvar"`
	StakeAuthority     string `json:"stakeAuthority"`
}

func NewParsedMerge(
	accounts []string,
) ParsedMerge {
	return ParsedMerge{
		Destination:        accounts[0],
		Source:             accounts[1],
		ClockSysvar:        accounts[2],
		StakeHistorySysvar: accounts[3],
		StakeAuthority:     accounts[4],
	}
}

//____________________________________________________________________________

type ParsedAuthorizeWithSeed struct {
	StakeAccount   string `json:"stakeAccount"`
	AuthorityBase  string `json:"authorBase"`
	NewAuthorized  string `json:"newAuthorized"`
	AuthorityType  string `json:"authorityType"`
	AuthoritySeed  string `json:"authoritySeed"`
	AuthorityOwner string `json:"authorityOwner"`

	ClockSysvar string `json:"clockSysvar,omitempty"`
	Custodian   string `json:"custodian,omitempty"`
}

func NewParsedAuthorizedWithSeed(
	accounts []string,
	instruction AuthorizeWithSeedInstruction,
) ParsedAuthorizeWithSeed {
	args := instruction.AuthorizeWithSeedArgs
	parsed := ParsedAuthorizeWithSeed{
		StakeAccount:   accounts[0],
		AuthorityBase:  accounts[1],
		NewAuthorized:  args.NewAuthorizedPubkey.String(),
		AuthoritySeed:  args.AuthoritySeed,
		AuthorityOwner: args.AuthorityOwner.String(),
	}
	if len(accounts) >= 3 {
		parsed.ClockSysvar = accounts[2]
	}
	if len(accounts) >= 4 {
		parsed.Custodian = accounts[3]
	}
	return parsed
}

//____________________________________________________________________________

type ParsedInitializeChecked struct {
	StakeAccount string `json:"stakeAccount"`
	RentSysvar   string `json:"rentSysvar"`
	Staker       string `json:"staker"`
	Withdrawer   string `json:"withdrawer"`
}

func NewInitializeChecked(
	accounts []string,
) ParsedInitializeChecked {
	return ParsedInitializeChecked{
		StakeAccount: accounts[0],
		RentSysvar:   accounts[1],
		Staker:       accounts[2],
		Withdrawer:   accounts[3],
	}
}

//____________________________________________________________________________

type ParsedAuthorizeChecked struct {
	StakeAccount  string `json:"stakeAccount"`
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
		StakeAccount:  accounts[0],
		ClockSysvar:   accounts[1],
		Authority:     accounts[2],
		NewAuthority:  accounts[3],
		AuthorityType: NewParsedStakeAuthorize(instruction.StakeAuthorize),
	}
}

//____________________________________________________________________________

type ParsedAuthorizeCheckedWithSeed struct {
	StakeAccount   string `json:"stakeAccount"`
	AuthorityBase  string `json:"authorityBase"`
	ClockSysvar    string `json:"clockSysvar"`
	NewAuthority   string `json:"newAuthority"`
	AuthorityType  string `json:"authorityType"`
	AuthoritySeed  string `json:"authoritySeed"`
	AuthorityOwner string `json:"authorityOwner"`

	Custodian string `json:"custodian,omitempty"`
}

func NewParsedAuthorizeCheckedWithSeed(
	accounts []string,
	instruction AuthorizeCheckedWithSeedInstruction,
) ParsedAuthorizeCheckedWithSeed {
	args := instruction.AuthorizeCheckedWithSeedArgs
	parsed := ParsedAuthorizeCheckedWithSeed{
		StakeAccount:   accounts[0],
		AuthorityBase:  accounts[1],
		ClockSysvar:    accounts[2],
		NewAuthority:   accounts[3],
		AuthorityType:  NewParsedStakeAuthorize(args.StakeAuthorize),
		AuthoritySeed:  args.AuthoritySeed,
		AuthorityOwner: args.AuthorityOwner.String(),
	}
	if len(accounts) >= 5 {
		parsed.Custodian = accounts[4]
	}
	return parsed
}

//____________________________________________________________________________

type ParsedSetLockupChecked struct {
	StakeAccount string       `json:"stakeAccount"`
	Custodian    string       `json:"custodian"`
	Lockup       ParsedLockup `json:"lockup"`
}

func NewParsedSetLockupChecked(
	accounts []string,
	instruction SetLockupCheckedInstruction,
) ParsedSetLockupChecked {
	parsed := ParsedSetLockupChecked{
		StakeAccount: accounts[0],
		Custodian:    accounts[1],
		Lockup:       NewParsedLockupFromCheckedArgs(instruction.LockupCheckedArgs),
	}
	if len(accounts) >= 3 {
		parsed.Lockup.Custodian = accounts[2]
	}
	return parsed
}

//____________________________________________________________________________

// Parsed instances used in parsed instructions

type ParsedAuthorized struct {
	Staker     string `json:"staker"`
	Withdrawer string `json:"withdrawer"`
}

func NewParsedAuthorized(authorized Authorized) ParsedAuthorized {
	return ParsedAuthorized{
		Staker:     authorized.Staker.String(),
		Withdrawer: authorized.Withdrawer.String(),
	}
}

func NewParsedStakeAuthorize(authorize StakeAuthorize) string {
	switch authorize {
	case Staker:
		return "staker"
	case Withdrawer:
		return "withdrawer"
	}
	return ""
}

type ParsedLockup struct {
	UnixTimestamp int64  `json:"unixTimestamp,omitempty"`
	Epoch         uint64 `json:"epoch,omitempty"`
	Custodian     string `json:"custodian,omitempty"`
}

func NewParsedLockup(lockup Lockup) ParsedLockup {
	return ParsedLockup{
		UnixTimestamp: lockup.UnixTimestamp,
		Epoch:         lockup.Epoch,
		Custodian:     lockup.Custodian.String(),
	}
}

func NewParsedLockupFromArgs(lockupArgs LockupArgs) ParsedLockup {
	lockup := ParsedLockup{}
	if lockupArgs.Custodian != nil {
		lockup.Custodian = lockupArgs.Custodian.String()
	}
	if lockupArgs.Epoch != nil {
		lockup.Epoch = *lockupArgs.Epoch
	}
	if lockupArgs.UnixTimestamp != nil {
		lockup.UnixTimestamp = *lockupArgs.UnixTimestamp
	}
	return lockup
}

func NewParsedLockupFromCheckedArgs(lockupArgs LockupCheckedArgs) ParsedLockup {
	lockup := ParsedLockup{}
	if lockupArgs.Epoch != nil {
		lockup.Epoch = *lockupArgs.Epoch
	}
	if lockupArgs.UnixTimestamp != nil {
		lockup.UnixTimestamp = *lockupArgs.UnixTimestamp
	}
	return lockup
}

//____________________________________________________________________________

type ParsedDeactivateDelinquent struct {
	StakeAccount string `json:"stakeAccount"`
	Delinquent   string `json:"delinquent"`
	Reference    string `json:"reference"`
}

func NewParsedDeactivateDelinquent(accounts []string) ParsedDeactivateDelinquent {
	return ParsedDeactivateDelinquent{
		StakeAccount: accounts[0],
		Delinquent:   accounts[1],
		Reference:    accounts[2],
	}
}
