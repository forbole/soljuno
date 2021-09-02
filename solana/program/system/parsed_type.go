package system

type ParsedCreateAccount struct {
	Source     string `json:"source"`
	NewAccount string `json:"newAccount"`
	Lamports   uint64 `json:"lamports"`
	Space      uint64 `json:"space"`
	Owner      string `json:"owner"`
}

func NewParsedCreateAccount(
	source,
	newAccount string,
	instruction CreateAccountInstruction,
) ParsedCreateAccount {
	return ParsedCreateAccount{
		Source:     source,
		NewAccount: newAccount,
		Lamports:   instruction.Lamports,
		Space:      instruction.Space,
		Owner:      instruction.Owner.String(),
	}
}

//____________________________________________________________________________

type ParsedAssign struct {
	Account string `json:"account"`
	Owner   string `json:"owner"`
}

func NewParsedAssign(
	account string,
	instruction AssignInstruction,
) ParsedAssign {
	return ParsedAssign{
		Account: account,
		Owner:   instruction.Owner.String(),
	}
}

//____________________________________________________________________________

type ParsedTransfer struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Lamports    uint64 `json:"lamports"`
}

func NewParsedTransfer(
	source,
	destination string,
	instruction TransferInstruction,
) ParsedTransfer {
	return ParsedTransfer{
		Source:      source,
		Destination: destination,
		Lamports:    instruction.Lamports,
	}
}

//____________________________________________________________________________

type ParsedCreateAccountWithSeed struct {
	Source     string `json:"source"`
	NewAccount string `json:"new_account"`
	Base       string `json:"base"`
	Seed       string `json:"seed"`
	Lamports   uint64 `json:"lamports"`
	Space      uint64 `json:"space"`
	Owner      string `json:"owner"`
}

func NewParsedCreateAccountWithSeed(
	source,
	newAccount string,
	instruction CreateAccountWithSeedInstruction,
) ParsedCreateAccountWithSeed {
	return ParsedCreateAccountWithSeed{
		Source:     source,
		NewAccount: newAccount,
		Base:       instruction.Base.String(),
		Seed:       instruction.Seed,
		Lamports:   instruction.Lamports,
		Owner:      instruction.Owner.String(),
	}
}

//____________________________________________________________________________

type ParsedAdvanceNonceAccount struct {
	NonceAccount            string `json:"nonceAccount"`
	RecentBlockHashesSysvar string `json:"recentBlockHashesSysvar"`
	NonceAuthority          string `json:"nonceAuthority"`
}

func NewParsedAdvanceNonceAccount(
	nonceAccount,
	recentBlockHashesSysvar,
	nonceAuthority string,
) ParsedAdvanceNonceAccount {
	return ParsedAdvanceNonceAccount{
		NonceAccount:            nonceAccount,
		RecentBlockHashesSysvar: recentBlockHashesSysvar,
		NonceAuthority:          nonceAuthority,
	}
}

//____________________________________________________________________________

type ParsedWithdrawNonceAccount struct {
	NonceAccount            string `json:"nonceAccount"`
	Destination             string `json:"destination"`
	RecentBlockHashesSysvar string `json:"recentBlockHashesSysvar"`
	RentSysvar              string `json:"rentSysvar"`
	NonceAuthority          string `json:"nonceAuthority"`
	Lamports                uint64 `json:"lamports"`
}

func NewParsedWithdrawNonceAccount(
	nonceAccount,
	destination,
	recentBlockHashesSysvar,
	rentSysvar,
	nonceAuthority string,
	instruction WithdrawNonceAccountInstruction,
) ParsedWithdrawNonceAccount {
	return ParsedWithdrawNonceAccount{
		NonceAccount:            nonceAccount,
		Destination:             destination,
		RecentBlockHashesSysvar: recentBlockHashesSysvar,
		RentSysvar:              rentSysvar,
		NonceAuthority:          nonceAuthority,
		Lamports:                instruction.Lamports,
	}
}

//____________________________________________________________________________

type ParsedInitializeNonceAccount struct {
	NonceAccount            string `json:"nonceAccount"`
	RecentBlockHashesSysvar string `json:"recentBlockHashesSysvar"`
	RentSysvar              string `json:"rentSysvar"`
	NonceAuthority          string `json:"nonceAuthority"`
}

func NewParsedInitializeNonceAccount(
	nonceAccount,
	recentBlockHashesSysvar,
	rentSysvar string,
	instruction InitializeNonceAccountInstruction,
) ParsedInitializeNonceAccount {
	return ParsedInitializeNonceAccount{
		NonceAccount:            nonceAccount,
		RecentBlockHashesSysvar: recentBlockHashesSysvar,
		RentSysvar:              rentSysvar,
		NonceAuthority:          instruction.Authority.String(),
	}
}

//____________________________________________________________________________

type ParsedAuthorizeNonceAccount struct {
	NonceAccount   string `json:"nonceAccount"`
	NonceAuthority string `json:"nonceAuthority"`
	NewAuthorized  string `json:"newAuthorized"`
}

func NewParsedAuthorizeNonceAccount(
	nonceAccount,
	nonceAuthority string,
	instruction AuthorizeNonceAccountInstruction,
) ParsedAuthorizeNonceAccount {
	return ParsedAuthorizeNonceAccount{
		NonceAccount:   nonceAccount,
		NonceAuthority: nonceAuthority,
		NewAuthorized:  instruction.Authority.String(),
	}
}

//____________________________________________________________________________

type ParsedAllocate struct {
	Account string `json:"account"`
	Space   uint64 `json:"space"`
}

func NewParsedAllocate(
	account string,
	instruction AllocateInstruction,
) ParsedAllocate {
	return ParsedAllocate{
		Account: account,
		Space:   instruction.Space,
	}
}

//____________________________________________________________________________

type ParsedAllocateWithSeed struct {
	Account string `json:"account"`
	Base    string `json:"base"`
	Seed    string `json:"seed"`
	Space   uint64 `json:"space"`
	Owner   string `json:"owner"`
}

func NewParsedAllocateWithSeed(
	account string,
	instruction AllocateWithSeedInstruction,
) ParsedAllocateWithSeed {
	return ParsedAllocateWithSeed{
		Account: account,
		Base:    instruction.Base.String(),
		Seed:    instruction.Seed,
		Space:   instruction.Space,
		Owner:   instruction.Owner.String(),
	}
}

//____________________________________________________________________________

type ParsedAssignWithSeed struct {
	Account string `json:"account"`
	Base    string `json:"base"`
	Seed    string `json:"seed"`
	Owner   string `json:"owner"`
}

func NewParsedAssignWithSeed(
	account string,
	instruction AssignWithSeedInstruction,
) ParsedAssignWithSeed {
	return ParsedAssignWithSeed{
		Account: account,
		Base:    instruction.Base.String(),
		Seed:    instruction.Seed,
		Owner:   instruction.Owner.String(),
	}
}

//____________________________________________________________________________

type ParsedTransferWithSeed struct {
	Source      string `json:"source"`
	SourceBase  string `json:"sourceBase"`
	Destination string `json:"destination"`
	Lamports    uint64 `json:"lamports"`
	SourceSeed  string `json:"sourceSeed"`
	SourceOwner string `json:"sourceOwner"`
}

func NewParsedTransferWithSeed(
	source,
	sourceBase,
	destination string,
	instruction TransferWithSeedInstruction,
) ParsedTransferWithSeed {
	return ParsedTransferWithSeed{
		Source:      source,
		SourceBase:  sourceBase,
		Destination: destination,
		Lamports:    instruction.Lamports,
		SourceSeed:  instruction.FromSeed,
		SourceOwner: instruction.FromOwner.String(),
	}
}
