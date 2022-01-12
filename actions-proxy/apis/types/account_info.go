package types

import (
	"encoding/base64"

	accountParser "github.com/forbole/soljuno/solana/account"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
)

type AccountInfoPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            AccountInfoArgs        `json:"input"`
}

type AccountInfoArgs struct {
	Address string `json:"address"`
}

type AccountInfoResponse struct {
	Data       [2]string        `json:"data"`
	Executable bool             `json:"executable"`
	Lamports   uint64           `json:"lamports"`
	Owner      string           `json:"program_owner"`
	RentEpoch  uint64           `json:"rent_epoch"`
	Parsed     *AccountResponse `json:"parsed"`
}

func NewAccountInfoResponse(info clienttypes.AccountInfo) (AccountInfoResponse, error) {
	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return AccountInfoResponse{}, err
	}
	parsed := accountParser.Parse(info.Value.Owner, bz)
	return AccountInfoResponse{
		Data:       info.Value.Data,
		Executable: info.Value.Executable,
		Lamports:   info.Value.Lamports,
		Owner:      info.Value.Owner,
		RentEpoch:  info.Value.RentEpoch,
		Parsed:     NewParsedAccountResponseFromInfo(parsed),
	}, nil
}

// --------------------------------

type AccountResponse struct {
	Type string            `json:"type"`
	Info ParsedAccountInfo `json:"info"`
}

func NewAccountResponse(typ string, info ParsedAccountInfo) *AccountResponse {
	return &AccountResponse{
		Type: typ,
		Info: info,
	}
}

type ParsedAccountInfo interface {
	IsParsedAccountInfo()
}

// --------------------------------

type NonceAccountInfo struct {
	BlockHash            string `json:"block_hash"`
	Authority            string `json:"authority"`
	LamportsPerSignature uint64 `json:"lamports_per_signature"`
}

func (a NonceAccountInfo) IsParsedAccountInfo() {}

// --------------------------------

type StakeAccountInfo struct {
	State string    `json:"state"`
	Meta  StakeMeta `json:"meta"`
	Stake *Stake    `json:"stake,omitempty"`
}

func (a StakeAccountInfo) IsParsedAccountInfo() {}

type StakeMeta struct {
	RentExemptReserve uint64          `json:"rent_exempt_reserve"`
	Authorized        StakeAuthorized `json:"authorized"`
	Lockup            StakeLockup     `json:"lockup"`
}

type StakeAuthorized struct {
	Staker     string `json:"staker"`
	Withdrawer string `json:"withdrawer"`
}

type StakeLockup struct {
	UnixTimestamp int64  `json:"unix_timestamp"`
	Epoch         uint64 `json:"epoch"`
	Custodian     string `json:"custodian"`
}

type Stake struct {
	Delegation      StakeDelegation `json:"delegation"`
	CreditsObserved uint64          `json:"credits_observed"`
}

type StakeDelegation struct {
	Voter              string  `json:"voter"`
	Stake              uint64  `json:"stake"`
	ActivationEpoch    uint64  `json:"activation_epoch"`
	DeactivationEpoch  uint64  `json:"deactivation_epoch"`
	WarmupCooldownRate float64 `json:"warmup_cooldown_rate"`
}

// --------------------------------

type TokenAccountInfo struct {
	Mint           string           `json:"token"`
	Owner          string           `json:"owner"`
	Amount         uint64           `json:"amount"`
	Delegate       *TokenDelegation `json:"delegate,omitempty"`
	CloseAuthority string           `json:"close_authority,omitempty"`
}

func (a TokenAccountInfo) IsParsedAccountInfo() {}

type TokenDelegation struct {
	Destination string `json:"destination"`
	Amount      uint64 `json:"amount"`
}

// --------------------------------

type TokenInfo struct {
	MintAuthority   string `json:"mint_authority"`
	Supply          uint64 `json:"supply"`
	Decimals        uint8  `json:"decimals"`
	FreezeAuthority string `json:"freeze_authority"`
}

func (t TokenInfo) IsParsedAccountInfo() {}

// -----------------------------

type VoteAccountInfo struct {
	Node       string `json:"node"`
	Withdrawer string `json:"withdrawer"`
	Commission uint8  `json:"commission"`
}

func (a VoteAccountInfo) IsParsedAccountInfo() {}

// -----------------------------

func NewParsedAccountResponseFromInfo(account interface{}) *AccountResponse {
	switch acc := account.(type) {
	case accountParser.NonceAccount:
		return NewAccountResponse(
			"nonce_account",
			NonceAccountInfo{
				BlockHash:            acc.BlockHash.String(),
				Authority:            acc.Authority.String(),
				LamportsPerSignature: acc.FeeCalculator.LamportsPerSignature,
			},
		)

	case accountParser.StakeAccount:
		meta := StakeMeta{
			RentExemptReserve: acc.Meta.RentExemptReserve,
			Authorized: StakeAuthorized{
				Staker:     acc.Meta.Authorized.Staker.String(),
				Withdrawer: acc.Meta.Authorized.Withdrawer.String(),
			},
			Lockup: StakeLockup{
				UnixTimestamp: acc.Meta.Lockup.UnixTimestamp,
				Epoch:         acc.Meta.Lockup.Epoch,
				Custodian:     acc.Meta.Lockup.Custodian.String(),
			},
		}
		var stake *Stake
		if acc.State.String() == "stake" {
			stake = &Stake{
				Delegation: StakeDelegation{
					Voter:              acc.Stake.Delegation.VoterPubkey.String(),
					Stake:              acc.Stake.Delegation.Stake,
					ActivationEpoch:    acc.Stake.Delegation.ActivationEpoch,
					DeactivationEpoch:  acc.Stake.Delegation.DeactivationEpoch,
					WarmupCooldownRate: acc.Stake.Delegation.WarmupCooldownRate,
				},
				CreditsObserved: acc.Stake.CreditsObserved,
			}
		}

		return NewAccountResponse(
			"stake_account",
			StakeAccountInfo{
				State: acc.State.String(),
				Meta:  meta,
				Stake: stake,
			})

	case accountParser.TokenAccount:
		var delegate *TokenDelegation
		if acc.Delegate.Option.Bool() {
			delegate = &TokenDelegation{
				Destination: acc.Delegate.String(),
				Amount:      acc.DelegateAmount,
			}
		}
		return NewAccountResponse(
			"token_account",
			TokenAccountInfo{
				Mint:           acc.Mint.String(),
				Owner:          acc.Owner.String(),
				Amount:         acc.Amount,
				Delegate:       delegate,
				CloseAuthority: acc.CloseAuthority.String(),
			})

	case accountParser.Token:
		return NewAccountResponse(
			"token",
			TokenInfo{
				MintAuthority:   acc.MintAuthority.String(),
				Supply:          acc.Supply,
				Decimals:        acc.Decimals,
				FreezeAuthority: acc.FreezeAuthority.String(),
			})

	case accountParser.VoteAccount:
		return NewAccountResponse(
			"vote_account",
			VoteAccountInfo{
				Node:       acc.Node.String(),
				Withdrawer: acc.Withdrawer.String(),
				Commission: acc.Commission,
			},
		)
	}
	return nil
}
