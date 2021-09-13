package token

import (
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/solana/parser"
	"github.com/forbole/soljuno/solana/types"
)

type Parser struct {
	parser.ProgramParser
}

func (Parser) Parse(accounts []string, data []byte) types.ParsedInstruction {
	decoder := bincode.NewDecoder()
	var id InstructionID
	decoder.Decode(data[:4], &id)
	switch id {
	case InitializeMint:
		var instruction InitializeMintInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"initializeMint",
			NewParsedParsedInitializeMint(
				accounts,
				instruction,
			),
		)

	case InitializeAccount:
		return types.NewParsedInstruction(
			"initializeAccount",
			NewParsedInitializeAccount(
				accounts,
			),
		)

	case InitializeMultisig:
		var instruction InitializeMultisigInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"initializeMultisig",
			NewParsedInitializeMultiisig(
				accounts,
				instruction,
			),
		)

	case Transfer:
		var instruction TransferInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"transfer",
			NewParsedTransfer(
				accounts,
				instruction,
			),
		)

	case Approve:
		var instruction ApproveInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"approve",
			NewParsedApprove(
				accounts,
				instruction,
			),
		)

	case Revoke:
		return types.NewParsedInstruction(
			"revoke",
			NewParsedRevoke(
				accounts,
			),
		)

	case SetAuthority:
		var instruction SetAuthorityInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"setAuthority",
			NewParsedSetAuthority(
				accounts,
				instruction,
			),
		)

	case MintTo:
		var instruction MintToInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"mintTo",
			NewParsedMintTo(
				accounts,
				instruction,
			),
		)

	case Burn:
		var instruction BurnInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"burn",
			NewParsedBurn(
				accounts,
				instruction,
			),
		)

	case CloseAccount:
		return types.NewParsedInstruction(
			"closeAccount",
			NewParsedCloseAccount(
				accounts,
			),
		)

	case FreezeAccount:
		return types.NewParsedInstruction(
			"freezeAccount",
			NewParsedFreezeAccount(
				accounts,
			),
		)

	case ThawAccount:
		return types.NewParsedInstruction(
			"thawAccount",
			NewParsedParsedThawAccount(
				accounts,
			),
		)

	case TransferChecked:
		var instruction TransferCheckedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"transferChecked",
			NewParsedTransferChecked(
				accounts,
				instruction,
			),
		)

	case ApproveChecked:
		var instruction ApproveCheckedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"approveChecked",
			NewParsedApproveChecked(
				accounts,
				instruction,
			),
		)

	case MintToChecked:
		var instruction MintToCheckedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"mintToChecked",
			NewParsedMintToChecked(
				accounts,
				instruction,
			),
		)

	case BurnChecked:
		var instruction BurnCheckedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"burnChecked",
			NewParsedBurnChecked(
				accounts,
				instruction,
			),
		)

	case InitializeAccount2:
		var instruction InitializeAccountInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"initializeAccount2",
			NewParsedInitializeAccount2(
				accounts,
				instruction,
			),
		)

	case SyncNative:
		return types.NewParsedInstruction(
			"syncNative",
			NewParsedSyncNative(
				accounts,
			),
		)

	case InitializeAccount3:
		var instruction InitializeAccountInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"initializeAccount3",
			NewParsedInitializeAccount3(
				accounts,
				instruction,
			),
		)

	case InitializeMultisig2:
		var instruction InitializeMultisigInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"initializeMultisig2",
			NewParsedInitializeMultiisig2(
				accounts,
				instruction,
			),
		)

	case InitializeMint2:
		var instruction InitializeMintInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"initializeMint2",
			NewParsedParsedInitializeMint2(
				accounts,
				instruction,
			),
		)

	}
	return types.NewParsedInstruction(
		"unknown",
		nil,
	)
}
