package snapshot

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	accountParser "github.com/forbole/soljuno/solana/account"
)

func ImportSnapshotCmd(cmdCfg *Config) *cobra.Command {
	return &cobra.Command{
		Use:     "import-snapshot [file]",
		Short:   "Import a snapshot at specific slot",
		PreRunE: ReadConfig(cmdCfg),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			context, err := GetParsingContext(cmdCfg)
			if err != nil {
				return err
			}
			if err != nil {
				return err
			}
			return StartImportSnapshot(context, args[0])
		},
	}
}

func StartImportSnapshot(ctx *Context, snapshotFile string) error {
	path, err := filepath.Abs(snapshotFile)
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() { err = file.Close() }()
	reader := bufio.NewReader(file)

	return handleSnapshot(ctx, reader)
}

func handleSnapshot(ctx *Context, reader *bufio.Reader) error {
	_, _, err := reader.ReadLine()
	if err != nil {
		return err
	}
	for {
		account, err := readSection(reader)
		if err == io.EOF {
			break
		}
		err = handleAccount(ctx, account.Pubkey)
		if err != nil {
			return err
		}
	}
	return nil
}

func readSection(reader *bufio.Reader) (Account, error) {
	var err error
	var detailBuf bytes.Buffer
	var l []byte
	count := 0
	var pubkey string
	for {
		l, _, err = reader.ReadLine()
		// If we're at the EOF, break.
		if err != nil {
			if err != io.EOF {
				return Account{}, err
			}
			break
		}

		if count == 0 {
			pubkey = string(l)[:len(l)-1]
			l = []byte(`account:`)
		}
		l = []byte(strings.Replace(string(l), "- ", "", 1))
		detailBuf.Write(l)
		count++
		if count == 7 {
			count = 0
			break
		}

		_, err = detailBuf.WriteString("\n")
		if err != nil {
			return Account{}, err
		}
	}
	if err == io.EOF {
		return Account{}, err
	}

	var account Account
	// Process the line here.
	err = yaml.Unmarshal(detailBuf.Bytes(), &account)
	if err != nil {
		return Account{}, err
	}
	account.Pubkey = pubkey
	return account, err
}

func handleAccount(ctx *Context, address string) error {
	info, err := ctx.Proxy.AccountInfo(address)
	if err != nil {
		return err
	}
	if info.Value == nil {
		return nil
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}
	account := accountParser.Parse(info.Value.Owner, bz)
	switch account := account.(type) {
	case accountParser.TokenMint:
		return updateToken(ctx, address, info.Context.Slot, account)

	case accountParser.TokenAccount:
		return updateTokenAccount(ctx, address, info.Context.Slot, account)

	case accountParser.Multisig:
		return updateMultisig(ctx, address, info.Context.Slot, account)

	case accountParser.StakeAccount:
		return updateStakeAccount(ctx, address, info.Context.Slot, account)

	case accountParser.VoteAccount:
		return updateVoteAccount(ctx, address, info.Context.Slot, account)

	case accountParser.NonceAccount:
		return updateNonceAccount(ctx, address, info.Context.Slot, account)

	case accountParser.BufferAccount:
		return updateBufferAccount(ctx, address, info.Context.Slot, account)

	case accountParser.ProgramAccount:
		return updateProgramAccount(ctx, address, info.Context.Slot, account)

	case accountParser.ProgramDataAccount:
		return updateProgramDataAccount(ctx, address, info.Context.Slot, account)

	case accountParser.StakeConfig:
		return nil

	case accountParser.ValidatorConfig:
		return updateValidatorConfig(ctx, address, info.Context.Slot, account)

	default:
		return nil
	}
}
