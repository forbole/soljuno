package stake

type ParsedInitialize struct {
	StakeAccount string           `json:"stakeAccount"`
	RentSysvar   string           `json:"rentSysvar"`
	Authorized   ParsedAuthorized `json:"authorized"`
	Lockup       ParsedLockup     `json:"lockup"`
}

//____________________________________________________________________________

type ParsedAuthorize struct {
	StakeAccount  string `json:"stakeAccount"`
	ClockSysvar   string `json:"clockSysvar"`
	Authority     string `json:"authority"`
	NewAuthority  string `json:"newAuthority"`
	AuthorityType string `json:"authorityType"`
	Custodian     string `json:"custodian,omitempty"`
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

//____________________________________________________________________________

type ParsedSplit struct {
	StakeAccount    string `json:"stakeAccount"`
	NewSplitAccount string `json:"newSplitAccount"`
	StakeAuthority  string `json:"stakeAuthority"`
	Lamports        uint64 `json:"lamports"`
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

//____________________________________________________________________________

type ParsedDeactivate struct {
	StakeAccount   string `json:"stakeAccount"`
	ClockSysvar    string `json:"clockSysvar"`
	StakeAuthority string `json:"stakeAuthority"`
}

//____________________________________________________________________________

type ParsedSetLockup struct {
	StakeAccount string `json:"stakeAccount"`
	Custodian    string `json:"custodian"`
	Lockup       ParsedLockup
}

//____________________________________________________________________________

type ParsedMerge struct {
	Destination        string `json:"destination"`
	Source             string `json:"source"`
	ClockSysvar        string `json:"clockSysvar"`
	StakeHistorySysvar string `json:"stakeHistorySysvar"`
	StakeAuthority     string `json:"stakeAuthority"`
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

//____________________________________________________________________________

type ParsedInitializeChecked struct {
	StakeAccount string `json:"stakeAccount"`
	RentSysvar   string `json:"rentSysvar"`
	Staker       string `json:"staker"`
	Withdrawer   string `json:"withdrawer"`
}

//____________________________________________________________________________

type ParsedAuthorizeChecked struct {
	StakeAccount  string `json:"stakeAccount"`
	ClockSysvar   string `json:"clockSysvar"`
	Authority     string `json:"authority"`
	NewAuthority  string `json:"newAuthority"`
	AuthorityType string `json:"authorityType"`
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

//____________________________________________________________________________

type ParsedSetLockupChecked struct {
	StakeAccount string       `json:"stakeAccount"`
	Custodian    string       `json:"custodian"`
	Lockup       ParsedLockup `json:"lockup"`
}

//____________________________________________________________________________

// Parsed instances used in parsed instructions

type ParsedAuthorized struct {
	Staker     string `json:"staker"`
	Withdrawer string `json:"withdrawer"`
}

type ParsedLockup struct {
	UnixTimestamp int64  `json:"unixTimestamp,omitempty"`
	Epoch         uint64 `json:"epoch,omitempty"`
	Custodian     string `json:"custodian,omitempty"`
}
