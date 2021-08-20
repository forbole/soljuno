package token

type AuthorityType uint32

const (
	// Authority to mint new tokens
	MintTokensType AuthorityType = iota

	// Authority to freeze any account associated with the Mint
	FreezeAccountType

	// Owner of a given token account
	AccountOwnerType

	// Authority to close a token account
	CloseAccountType
)
