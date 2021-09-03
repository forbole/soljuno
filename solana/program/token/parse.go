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
			"unknown",
			NewParsedParsedInitializeMint(
				accounts,
				instruction,
			),
		)

	case InitializeAccount:
		return types.NewParsedInstruction(
			"unknown",
			NewParsedInitializeAccount(
				accounts,
			),
		)

	case InitializeMultisig:
		var instruction InitializeMultisigInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"unknown",
			NewParsedInitializeMultiisig(
				accounts,
				instruction,
			),
		)

	case Transfer:
		var instruction TransferInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"unknown",
			NewParsedTransfer(
				accounts,
				instruction,
			),
		)

	case Approve:
		var instruction ApproveInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"unknown",
			NewParsedApprove(
				accounts,
				instruction,
			),
		)

	case Revoke:
		return types.NewParsedInstruction(
			"unknown",
			NewParsedRevoke(
				accounts,
			),
		)

	case SetAuthority:
		var instruction SetAuthorityInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"unknown",
			NewParsedSetAuthority(
				accounts,
				instruction,
			),
		)

	case MintTo:
		var instruction MintToInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"unknown",
			NewParsedMintTo(
				accounts,
				instruction,
			),
		)

	case Burn:
		var instruction BurnInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"unknown",
			NewParsedBurn(
				accounts,
				instruction,
			),
		)

	case CloseAccount:
		return types.NewParsedInstruction(
			"unknown",
			NewParsedCloseAccount(
				accounts,
			),
		)

	case FreezeAccount:
		return types.NewParsedInstruction(
			"unknown",
			NewParsedFreezeAccount(
				accounts,
			),
		)

	case ThawAccount:
		return types.NewParsedInstruction(
			"unknown",
			NewParsedParsedThawAccount(
				accounts,
			),
		)

	case TransferChecked:
		var instruction TransferCheckedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"unknown",
			NewParsedTransferChecked(
				accounts,
				instruction,
			),
		)

	case ApproveChecked:
		var instruction ApproveCheckedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"unknown",
			NewParsedApproveChecked(
				accounts,
				instruction,
			),
		)

	case MintToChecked:
		var instruction MintToCheckedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"unknown",
			NewParsedMintToChecked(
				accounts,
				instruction,
			),
		)

	case BurnChecked:
		var instruction BurnCheckedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"unknown",
			NewParsedBurnChecked(
				accounts,
				instruction,
			),
		)

	case InitializeAccount2:
		var instruction InitializeAccountInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"unknown",
			NewParsedInitializeAccount2(
				accounts,
				instruction,
			),
		)

	case SyncNative:
		return types.NewParsedInstruction(
			"unknown",
			NewParsedSyncNative(
				accounts,
			),
		)

	case InitializeAccount3:
		var instruction InitializeAccountInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"unknown",
			NewParsedInitializeAccount3(
				accounts,
				instruction,
			),
		)

	case InitializeMultisig2:
		var instruction InitializeMultisigInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"unknown",
			NewParsedInitializeMultiisig2(
				accounts,
				instruction,
			),
		)

	case InitializeMint2:
		var instruction InitializeMintInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"unknown",
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
