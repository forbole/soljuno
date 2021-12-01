package tokenswap

type InstructionID uint8

const (
	// Initializes a new swap
	Initialize InstructionID = iota

	// Swap the tokens in the pool
	Swap

	// Deposit both types of tokens into the pool.  The output is a "pool"
	// token representing ownership in the pool. Inputs are converted to
	// the current ratio.
	DepositAllTokenTypes

	// Withdraw both types of tokens from the pool at the current ratio, given
	// pool tokens.  The pool tokens are burned in exchange for an equivalent
	// amount of token A and B.
	WithdrawAllTokenTypes

	// Deposit one type of tokens into the pool.  The output is a "pool" token
	// representing ownership into the pool. Input token is converted as if
	// a swap and deposit all token types were performed.
	DepositSingleTokenTypeExactAmountIn

	// Withdraw one token type from the pool at the current ratio given the
	// exact amount out expected.
	WithdrawSingleTokenTypeExactAmountOut
)
