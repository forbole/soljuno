package token

import (
	"math"
	"strconv"

	"github.com/forbole/soljuno/solana/client/types"
)

type ParsedInitializeMint struct {
	Mint          string `json:"mint"`
	Decimals      uint8  `json:"decimals"`
	MintAuthority string `json:"mintAuthority"`

	RentSysvar      string `json:"rentSysvar,omitempty"`
	FreezeAuthority string `json:"freezeAuthority,omitempty"`
}

func NewParsedParsedInitializeMint(
	accounts []string,
	instruction InitializeMintInstruction,
) ParsedInitializeMint {
	parsed := ParsedInitializeMint{
		Mint:          accounts[0],
		Decimals:      instruction.Decimals,
		MintAuthority: instruction.MintAuthority.String(),
		RentSysvar:    accounts[1],
	}

	if instruction.FreezeAuthority != nil {
		parsed.FreezeAuthority = instruction.FreezeAuthority.String()
	}
	return parsed
}

func NewParsedParsedInitializeMint2(
	accounts []string,
	instruction InitializeMintInstruction,
) ParsedInitializeMint {
	parsed := ParsedInitializeMint{
		Mint:          accounts[0],
		Decimals:      instruction.Decimals,
		MintAuthority: instruction.MintAuthority.String(),
	}

	if instruction.FreezeAuthority != nil {
		parsed.FreezeAuthority = instruction.FreezeAuthority.String()
	}
	return parsed
}

//____________________________________________________________________________

type ParsedInitializeAccount struct {
	Account    string `json:"account"`
	Mint       string `json:"mint"`
	Owner      string `json:"owner"`
	RentSysvar string `json:"rentSysvar,omitempty"`
}

func NewParsedInitializeAccount(
	accounts []string,
) ParsedInitializeAccount {
	return ParsedInitializeAccount{
		Account:    accounts[0],
		Mint:       accounts[1],
		Owner:      accounts[2],
		RentSysvar: accounts[3],
	}
}

func NewParsedInitializeAccount2(
	accounts []string,
	instruction InitializeAccountInstruction,
) ParsedInitializeAccount {
	return ParsedInitializeAccount{
		Account:    accounts[0],
		Mint:       accounts[1],
		Owner:      instruction.Owner.String(),
		RentSysvar: accounts[2],
	}
}

func NewParsedInitializeAccount3(
	accounts []string,
	instruction InitializeAccountInstruction,
) ParsedInitializeAccount {
	return ParsedInitializeAccount{
		Account: accounts[0],
		Mint:    accounts[1],
		Owner:   instruction.Owner.String(),
	}
}

//____________________________________________________________________________

type ParsedInitializeMultisig struct {
	MultiSig   string   `json:"multiSig"`
	RentSysvar string   `json:"rentSysvar,omitempty"`
	Signers    []string `json:"signers"`
	M          uint8    `json:"m"`
}

func NewParsedInitializeMultiisig(
	accounts []string,
	instruction InitializeMultisigInstruction,
) ParsedInitializeMultisig {
	signers := accounts[2:]
	return ParsedInitializeMultisig{
		MultiSig:   accounts[0],
		RentSysvar: accounts[1],
		Signers:    signers,
		M:          instruction.Amount,
	}
}

func NewParsedInitializeMultiisig2(
	accounts []string,
	instruction InitializeMultisigInstruction,
) ParsedInitializeMultisig {
	signers := accounts[2:]
	return ParsedInitializeMultisig{
		MultiSig: accounts[0],
		Signers:  signers,
		M:        instruction.Amount,
	}
}

//____________________________________________________________________________

type ParsedTransfer struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Amount      string `json:"amount"`

	Authority         string   `json:"authority,omitempty"`
	MultisigAuthority string   `json:"multisig,omitempty"`
	Signers           []string `json:"signers,omitempty"`
}

func NewParsedTransfer(
	accounts []string,
	instruction TransferInstruction,
) ParsedTransfer {
	parsed := ParsedTransfer{
		Source:      accounts[0],
		Destination: accounts[1],
		Amount:      strconv.FormatUint(instruction.Amount, 10),
	}

	if len(accounts) > 3 {
		parsed.MultisigAuthority = accounts[3]
		parsed.Signers = accounts[4:]
	} else {
		parsed.Authority = accounts[3]
	}
	return parsed
}

//____________________________________________________________________________

type ParsedApprove struct {
	Source   string `json:"source"`
	Delegate string `json:"delegate"`
	Amount   string `json:"amount"`

	Owner         string   `json:"owner,omitempty"`
	MultisigOwner string   `json:"multisigOwner,omitempty"`
	Signers       []string `json:"signers,omitempty"`
}

func NewParsedApprove(
	accounts []string,
	instruction ApproveInstruction,
) ParsedApprove {
	parsed := ParsedApprove{
		Source:   accounts[0],
		Delegate: accounts[1],
		Amount:   strconv.FormatUint(instruction.Amount, 10),
	}

	if len(accounts) > 3 {
		parsed.MultisigOwner = accounts[2]
		parsed.Signers = accounts[2:]
	} else {
		parsed.Owner = accounts[2]
	}
	return parsed
}

//____________________________________________________________________________

type ParsedRevoke struct {
	Source string `json:"source"`

	Owner         string   `json:"owner,omitempty"`
	MultisigOwner string   `json:"multisigOwner,omitempty"`
	Signers       []string `json:"signers,omitempty"`
}

func NewParsedRevoke(
	accounts []string,
) ParsedRevoke {
	parsed := ParsedRevoke{Source: accounts[0]}

	if len(accounts) > 2 {
		parsed.MultisigOwner = accounts[1]
		parsed.Signers = accounts[2:]
	} else {
		parsed.Owner = accounts[1]
	}

	return parsed
}

//____________________________________________________________________________

type ParsedSetAuthority struct {
	Mint    string `json:"mint,omitempty"`
	Account string `json:"account,omitempty"`

	AuthorityType string `json:"authorityType"`
	NewAuthority  string `json:"newAuthority"`

	Authority         string   `json:"authority,omitempty"`
	MultisigAuthority string   `json:"multisigAuthority"`
	Signers           []string `json:"signers,omitempty"`
}

func NewParsedSetAuthority(
	accounts []string,
	instruction SetAuthorityInstruction,
) ParsedSetAuthority {
	parsed := ParsedSetAuthority{}
	if instruction.AuthorityType == MintTokensType || instruction.AuthorityType == FreezeAccountType {
		parsed.Mint = accounts[0]
	}
	if instruction.AuthorityType == AccountOwnerType || instruction.AuthorityType == CloseAccountType {
		parsed.Account = accounts[0]
	}
	if instruction.NewAuthority != nil {
		parsed.NewAuthority = instruction.NewAuthority.String()
	}

	if len(accounts) > 2 {
		parsed.MultisigAuthority = accounts[1]
		parsed.Signers = accounts[2:]
	} else {
		parsed.Authority = accounts[1]
	}

	return parsed
}

//____________________________________________________________________________

type ParsedMintTo struct {
	Mint    string `json:"mint"`
	Account string `json:"account"`
	Amount  string `json:"amount"`

	MintAuthority     string   `json:"mintAuthority,omitempty"`
	MultisigAuthority string   `json:"multisigAuthority,omitempty"`
	Signers           []string `json:"signers,omitempty"`
}

func NewParsedMintTo(
	accounts []string,
	instruction MintToInstruction,
) ParsedMintTo {
	parsed := ParsedMintTo{
		Mint:    accounts[0],
		Account: accounts[1],
		Amount:  strconv.FormatUint(instruction.Amount, 10),
	}

	if len(accounts) > 3 {
		parsed.MultisigAuthority = accounts[2]
		parsed.Signers = accounts[3:]
	} else {
		parsed.MintAuthority = accounts[2]
	}
	return parsed
}

//____________________________________________________________________________

type ParsedBurn struct {
	Account string `json:"account"`
	Mint    string `json:"mint"`
	Amount  string `json:"amount"`

	Authority         string   `json:"authority,omitempty"`
	MultisigAuthority string   `json:"multisigAuthority,omitempty"`
	Signers           []string `json:"singers,omitempty"`
}

func NewParsedBurn(
	accounts []string,
	instruction BurnInstruction,
) ParsedBurn {
	parsed := ParsedBurn{
		Account: accounts[0],
		Mint:    accounts[1],
		Amount:  strconv.FormatUint(instruction.Amount, 10),
	}

	if len(accounts) > 3 {
		parsed.MultisigAuthority = accounts[2]
		parsed.Signers = accounts[3:]
	} else {
		parsed.Authority = accounts[2]
	}
	return parsed
}

//____________________________________________________________________________

type ParsedCloseAccount struct {
	Account     string `json:"account"`
	Destination string `json:"destination"`

	Owner             string   `json:"owner,omitempty"`
	MultisigAuthority string   `json:"multisigAuthority,omitempty"`
	Signers           []string `json:"singers,omitempty"`
}

func NewParsedCloseAccount(
	accounts []string,
) ParsedCloseAccount {
	parsed := ParsedCloseAccount{
		Account:     accounts[0],
		Destination: accounts[1],
	}

	if len(accounts) > 3 {
		parsed.MultisigAuthority = accounts[2]
		parsed.Signers = accounts[3:]
	} else {
		parsed.Owner = accounts[2]
	}
	return parsed
}

//____________________________________________________________________________

type ParsedFreezeAccount struct {
	Account string `json:"account"`
	Mint    string `json:"mint"`

	FreezeAuthority         string   `json:"freezeAuthority,omitempty"`
	MultisigFreezeAuthority string   `json:"multisigFreezeAuthority,omitempty"`
	Signers                 []string `json:"singers,omitempty"`
}

func NewParsedFreezeAccount(
	accounts []string,
) ParsedFreezeAccount {
	parsed := ParsedFreezeAccount{
		Account: accounts[0],
		Mint:    accounts[1],
	}

	if len(accounts) > 3 {
		parsed.MultisigFreezeAuthority = accounts[2]
		parsed.Signers = accounts[3:]
	} else {
		parsed.FreezeAuthority = accounts[2]
	}
	return parsed
}

//____________________________________________________________________________

type ParsedThawAccount struct {
	Account string `json:"account"`
	Mint    string `json:"mint"`

	FreezeAuthority         string   `json:"freezeAuthority,omitempty"`
	MultisigFreezeAuthority string   `json:"multisigFreezeAuthority,omitempty"`
	Signers                 []string `json:"singers,omitempty"`
}

func NewParsedParsedThawAccount(
	accounts []string,
) ParsedThawAccount {
	parsed := ParsedThawAccount{
		Account: accounts[0],
		Mint:    accounts[1],
	}

	if len(accounts) > 3 {
		parsed.MultisigFreezeAuthority = accounts[2]
		parsed.Signers = accounts[3:]
	} else {
		parsed.FreezeAuthority = accounts[2]
	}
	return parsed
}

//____________________________________________________________________________

type ParsedTransferChecked struct {
	Source      string              `json:"source"`
	Mint        string              `json:"mint"`
	Destination string              `json:"destination"`
	TokenAmount types.UiTokenAmount `json:"tokenAmount"`

	Authority         string   `json:"authority,omitempty"`
	MultisigAuthority string   `json:"multisigAuthority,omitempty"`
	Signers           []string `json:"signers,omitempty"`
}

func NewParsedTransferChecked(
	accounts []string,
	instruction TransferCheckedInstruction,
) ParsedTransferChecked {
	parsed := ParsedTransferChecked{
		Source:      accounts[0],
		Mint:        accounts[1],
		Destination: accounts[2],
		TokenAmount: TokenAmountToUiAmount(
			instruction.Amount,
			instruction.Decimals,
		),
	}

	if len(accounts) > 4 {
		parsed.MultisigAuthority = accounts[3]
		parsed.Signers = accounts[4:]
	} else {
		parsed.Authority = accounts[3]
	}
	return parsed
}

//____________________________________________________________________________

type ParsedApproveChecked struct {
	Source      string              `json:"source"`
	Mint        string              `json:"mint"`
	Delegate    string              `json:"delegate"`
	TokenAmount types.UiTokenAmount `json:"tokenAmount"`

	Owner         string   `json:"owner,omitempty"`
	MultisigOwner string   `json:"multisigOwner,omitempty"`
	Signers       []string `json:"singers,omitempty"`
}

func NewParsedApproveChecked(
	accounts []string,
	instruction ApproveCheckedInstruction,
) ParsedApproveChecked {
	parsed := ParsedApproveChecked{
		Source:   accounts[0],
		Mint:     accounts[1],
		Delegate: accounts[2],
		TokenAmount: TokenAmountToUiAmount(
			instruction.Amount,
			instruction.Decimals,
		),
	}

	if len(accounts) > 4 {
		parsed.MultisigOwner = accounts[3]
		parsed.Signers = accounts[4:]
	} else {
		parsed.Owner = accounts[3]
	}
	return parsed
}

//____________________________________________________________________________

type ParsedMintToChecked struct {
	Mint        string              `json:"mint"`
	Account     string              `json:"account"`
	TokenAmount types.UiTokenAmount `json:"tokenAmount"`

	MintAuthority     string   `json:"mintAuthority,omitempty"`
	MultisigAuthority string   `json:"multisigAuthority,omitempty"`
	Signers           []string `json:"signers,omitempty"`
}

func NewParsedMintToChecked(
	accounts []string,
	instruction MintToCheckedInstruction,
) ParsedMintToChecked {
	parsed := ParsedMintToChecked{
		Mint:    accounts[0],
		Account: accounts[1],
		TokenAmount: TokenAmountToUiAmount(
			instruction.Amount,
			instruction.Decimals,
		),
	}

	if len(accounts) > 3 {
		parsed.MultisigAuthority = accounts[2]
		parsed.Signers = accounts[3:]
	} else {
		parsed.MintAuthority = accounts[2]
	}
	return parsed
}

//____________________________________________________________________________

type ParsedBurnChecked struct {
	Account     string              `json:"account"`
	Mint        string              `json:"mint"`
	TokenAmount types.UiTokenAmount `json:"tokenAmount"`

	Authority         string   `json:"authority,omitempty"`
	MultisigAuthority string   `json:"multisigAuthority,omitempty"`
	Signers           []string `json:"singers,omitempty"`
}

func NewParsedBurnChecked(
	accounts []string,
	instruction BurnCheckedInstruction,
) ParsedBurnChecked {
	parsed := ParsedBurnChecked{
		Account: accounts[0],
		Mint:    accounts[1],
		TokenAmount: TokenAmountToUiAmount(
			instruction.Amount,
			instruction.Decimals,
		),
	}

	if len(accounts) > 3 {
		parsed.MultisigAuthority = accounts[2]
		parsed.Signers = accounts[3:]
	} else {
		parsed.Authority = accounts[2]
	}
	return parsed
}

//____________________________________________________________________________

type ParsedSyncNative struct {
	Account string `json:"account"`
}

func NewParsedSyncNative(
	accounts []string,
) ParsedSyncNative {
	return ParsedSyncNative{
		Account: accounts[0],
	}
}

//____________________________________________________________________________

func TokenAmountToUiAmount(
	amount uint64,
	decimals uint8,
) types.UiTokenAmount {
	amountDecimals := float64(amount)
	if decimals != 0 {
		amountDecimals /= math.Pow(10, float64(decimals))
	}
	return types.UiTokenAmount{
		UiAmount: amountDecimals,
		Amount:   strconv.FormatUint(amount, 10),
		Decimals: decimals,
	}
}
