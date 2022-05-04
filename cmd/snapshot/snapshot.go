package snapshot

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	cmdtypes "github.com/forbole/soljuno/cmd/types"
	"github.com/forbole/soljuno/solana/account/parser"
)

const (
	FlagParallelize = "parallelize"
	FlagSkip        = "skip"
)

func ImportSnapshotCmd(cmdCfg *cmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "import-snapshot [file]",
		Short:   "Import a snapshot at specific slot",
		PreRunE: cmdtypes.ReadConfig(cmdCfg),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			context, err := GetSnapshotContext(cmdCfg)
			if err != nil {
				return err
			}
			parallelize, err := cmd.Flags().GetInt(FlagParallelize)
			if err != nil {
				return err
			}

			skip, err := cmd.Flags().GetInt(FlagSkip)
			if err != nil {
				return err
			}

			return StartImportSnapshot(context, args[0], skip, parallelize)
		},
	}
	cmd.Flags().Int(FlagParallelize, 100, "the amount of accounts to process at a time")
	cmd.Flags().Int(FlagSkip, 0, "the amount of accounts to skip")
	return cmd
}

func StartImportSnapshot(ctx *Context, snapshotFile string, skip int, parallelize int) error {
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

	return handleSnapshot(ctx, reader, skip, parallelize)
}

// handleSnapshot handles all accounts inside the snapshot file
func handleSnapshot(ctx *Context, reader *bufio.Reader, skip int, parallelize int) error {
	wg := new(sync.WaitGroup)
	for i := 0; ; i++ {
		if skip > i {
			continue
		}

		// Sleep when pool is full or reach the parallelize limit
		if ctx.Pool.Free() == 0 || (i+1)%parallelize == 0 {
			time.Sleep(time.Second)
		}

		// Read account section from yaml
		pubkey, bz, err := readSection(reader)
		// Break the loop when there is no new account
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		var account Account
		err = yaml.Unmarshal(bz.Bytes(), &account)
		if err != nil {
			return err
		}
		account.Pubkey = pubkey

		wg.Add(1)
		err = ctx.Pool.Submit(
			func() {
				defer wg.Done()
				ctx.Logger.Info("Start handling account", "address", pubkey, "index", i)
				delay := 0
				for {
					err := handleAccount(ctx, account)
					if err != nil {
						ctx.Logger.Error("failed to import account", "address", pubkey, "err", err)
						ctx.Logger.Info("retry to import account", "address", pubkey)
						if delay <= 100 {
							delay += 3
						}
						time.Sleep(time.Duration(delay) * time.Second)
					} else {
						return
					}
				}
			},
		)
		if err != nil {
			return err
		}
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
			pubkey = string(l)
			l = []byte(`account:`)
		}
		_, err := buf.Write(l)
		if err != nil {
			return "", bytes.Buffer{}, err
		}

		_, err = buf.WriteString("\n")
		if err != nil {
			return "", bytes.Buffer{}, err
		}
	}
	// Check if it is the last line
	if err == io.EOF {
		return "", bytes.Buffer{}, err
	}
	return pubkey, buf, nil
}

func handleAccount(ctx *Context, account Account) (err error) {
	defer func() {
		r := recover()
		if r != nil {
			ctx.Logger.Error("failed to handle account", "address", account.Pubkey, "err", err)
		}
	}()

	// Update account data from node
	address, balance, err := account.ToBalance()
	if err != nil {
		return err
	}
	err = ctx.Database.SaveAccountBalances(account.Detail.Slot, []string{address}, []uint64{balance})
	if err != nil {
		return err
	}

	info, err := ctx.Proxy.GetAccountInfo(address)
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

	switch account := parser.Parse(info.Value.Owner, bz).(type) {
	case parser.Token:
		return saveToken(ctx.Database, address, info.Context.Slot, account)

	case parser.TokenAccount:
		if err := saveTokenAccount(ctx.Database, address, info.Context.Slot, account); err != nil {
			return err
		}
		return saveTokenBalance(ctx.Database, address, info.Context.Slot, account)

	case parser.Multisig:
		return saveMultisig(ctx.Database, address, info.Context.Slot, account)

	case parser.StakeAccount:
		return updateStakeAccount(ctx.Database, address, info.Context.Slot, account)

	case parser.VoteAccount:
		return saveVoteAccount(ctx.Database, address, info.Context.Slot, account)

	case parser.NonceAccount:
		return updateNonceAccount(ctx.Database, address, info.Context.Slot, account)

	case parser.BufferAccount:
		return updateBufferAccount(ctx.Database, address, info.Context.Slot, account)

	case parser.ProgramAccount:
		return updateProgramAccount(ctx.Database, address, info.Context.Slot, account)

	case parser.ProgramDataAccount:
		return updateProgramDataAccount(ctx.Database, address, info.Context.Slot, account)

	case parser.ValidatorConfig:
		return updateValidatorConfig(ctx.Database, address, info.Context.Slot, account)

	default:
		return nil
	}
}
