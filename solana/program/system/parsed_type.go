package system

type ParsedCreateAccount struct {
	Source     string `json:"source"`
	NewAccount string `json:"newAccount"`
	Lamports   uint64 `json:"lamports"`
	Space      uint64 `json:"space"`
	Owner      string `json:"owner"`
}

func NewParsedCreateAccount(
	accounts []string,
	instruction CreateAccountInstruction,
) ParsedCreateAccount {
	return ParsedCreateAccount{
		Source:     accounts[0],
		NewAccount: accounts[1],
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
	accounts []string,
	instruction AssignInstruction,
) ParsedAssign {
	return ParsedAssign{
		Account: accounts[0],
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
	accounts []string,
	instruction TransferInstruction,
) ParsedTransfer {
	return ParsedTransfer{
		Source:      accounts[0],
		Destination: accounts[1],
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
	accounts []string,
	instruction CreateAccountWithSeedInstruction,
) ParsedCreateAccountWithSeed {
	return ParsedCreateAccountWithSeed{
		Source:     accounts[0],
		NewAccount: accounts[1],
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
	accounts []string,
) ParsedAdvanceNonceAccount {
	return ParsedAdvanceNonceAccount{
		NonceAccount:            accounts[0],
		RecentBlockHashesSysvar: accounts[1],
		NonceAuthority:          accounts[2],
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
	accounts []string,
	instruction WithdrawNonceAccountInstruction,
) ParsedWithdrawNonceAccount {
	return ParsedWithdrawNonceAccount{
		NonceAccount:            accounts[0],
		Destination:             accounts[1],
		RecentBlockHashesSysvar: accounts[2],
		RentSysvar:              accounts[3],
		NonceAuthority:          accounts[4],
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
	accounts []string,
	instruction InitializeNonceAccountInstruction,
) ParsedInitializeNonceAccount {
	return ParsedInitializeNonceAccount{
		NonceAccount:            accounts[0],
		RecentBlockHashesSysvar: accounts[1],
		RentSysvar:              accounts[2],
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
	accounts []string,
	instruction AuthorizeNonceAccountInstruction,
) ParsedAuthorizeNonceAccount {
	return ParsedAuthorizeNonceAccount{
		NonceAccount:   accounts[0],
		NonceAuthority: accounts[1],
		NewAuthorized:  instruction.Authority.String(),
	}
}

//____________________________________________________________________________

type ParsedAllocate struct {
	Account string `json:"account"`
	Space   uint64 `json:"space"`
}

func NewParsedAllocate(
	accounts []string,
	instruction AllocateInstruction,
) ParsedAllocate {
	return ParsedAllocate{
		Account: accounts[0],
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
	accounts []string,
	instruction AllocateWithSeedInstruction,
) ParsedAllocateWithSeed {
	return ParsedAllocateWithSeed{
		Account: accounts[0],
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
	accounts []string,
	instruction AssignWithSeedInstruction,
) ParsedAssignWithSeed {
	return ParsedAssignWithSeed{
		Account: accounts[0],
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
	accounts []string,
	instruction TransferWithSeedInstruction,
) ParsedTransferWithSeed {
	return ParsedTransferWithSeed{
		Source:      accounts[0],
		SourceBase:  accounts[1],
		Destination: accounts[2],
		Lamports:    instruction.Lamports,
		SourceSeed:  instruction.FromSeed,
		SourceOwner: instruction.FromOwner.String(),
	}
}
