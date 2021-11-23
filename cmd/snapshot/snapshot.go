package snapshot

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

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
			context, err := GetSnapshotContext(cmdCfg)
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
	defer func() { _ = file.Close() }()
	reader := bufio.NewReader(file)

	return handleSnapshot(ctx, reader)
}

func handleSnapshot(ctx *Context, reader *bufio.Reader) error {
	_, _, err := reader.ReadLine()
	if err != nil {
		return err
	}
	wg := new(sync.WaitGroup)
	for {
		if ctx.Pool.Free() == 0 {
			time.Sleep(1 * time.Second)
			continue
		}
		pubkey, buf, err := readSection(reader)
		if err == io.EOF {
			break
		}

		var account Account
		// Process the line here.
		err = yaml.Unmarshal(buf.Bytes(), &account)
		if err != nil {
			return err
		}
		account.Pubkey = pubkey
		wg.Add(1)
		err = ctx.Pool.Submit(
			func() {
				defer wg.Done()
				ctx.Logger.Info("Start handling account", "address", pubkey)
				err = handleAccount(ctx, account.Pubkey)
				if err != nil {
					ctx.Logger.Error("failed to import account", "address", pubkey, "err", err)
				}
			},
		)
	}
	wg.Wait()
	return nil
}

// readSection reads a section of account in the snapshot returns a pubkey and a buffer of detail
func readSection(reader *bufio.Reader) (string, bytes.Buffer, error) {
	var err error
	var buf bytes.Buffer
	var pubkey string
	for count := 0; count < 7; count++ {
		var l []byte
		l, _, err = reader.ReadLine()
		// If we're at the EOF, break.
		if err == io.EOF {
			break
		} else if err != nil {
			return "", bytes.Buffer{}, err
		}

		if count == 0 {
			pubkey = string(l)[:len(l)-1]
			l = []byte(`account:`)
		}
		l = []byte(strings.Replace(string(l), "- ", "", 1))
		buf.Write(l)

		_, err = buf.WriteString("\n")
		if err != nil {
			return "", bytes.Buffer{}, err
		}
	}
	if err == io.EOF {
		return "", bytes.Buffer{}, err
	}
	return pubkey, buf, nil
}

func handleAccount(ctx *Context, address string) error {
	info, err := ctx.Proxy.AccountInfo(address)
	if err != nil {
		return err
	}
	if info.Value == nil {
		return nil
	}
	err = updateAccountBalance(ctx, address, info)
	if err != nil {
		return err
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}
	account := accountParser.Parse(info.Value.Owner, bz)
	switch account := account.(type) {
	case accountParser.Token:
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

	case accountParser.ValidatorConfig:
		return updateValidatorConfig(ctx, address, info.Context.Slot, account)

	default:
		return nil
	}
}
