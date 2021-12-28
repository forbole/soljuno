package tokenlist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	cmdtypes "github.com/forbole/soljuno/cmd/types"
	dbtypes "github.com/forbole/soljuno/db/types"
)

func ImportTokenListCmd(cmdCfg *cmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "import-tokenlist [file]",
		Short:   "Import a tokenlist to the token unit table",
		PreRunE: cmdtypes.ReadConfig(cmdCfg),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			context, err := GetTokenListContext(cmdCfg)
			if err != nil {
				return err
			}
			return ImportTokenList(context, args[0])
		},
	}
	return cmd
}

func ImportTokenList(ctx *Context, file string) error {
	tokenList, err := getTokenList(file)
	if err != nil {
		return err
	}
	tokens := getMainnetTokens(tokenList)
	ctx.Logger.Info(fmt.Sprintf("Importing %d tokens into database...", len(tokens)))
	count := 0
	var rows []dbtypes.TokenUnitRow
	// Add solana to the list
	rows = append(rows, dbtypes.NewTokenUnitRow("", "solana", "Solana", "", tokenList.LogoURI, "https://solana.com"))

	tokenMap := make(map[string]bool)
	for _, token := range tokens {
		if len(rows) >= 1000 {
			err := ctx.Database.SaveTokenUnits(rows)
			if err != nil {
				return err
			}
			rows = nil
			count = 0
		}
		rows = append(rows, dbtypes.NewTokenUnitRow(
			token.Address,
			token.Extensions.CoingeckoID,
			token.Name,
			token.LogoURI,
			token.Extensions.Description,
			token.Extensions.Website,
		))
		tokenMap[token.Address] = true
		count++
	}
	return ctx.Database.SaveTokenUnits(rows)
}

func getTokenList(listFile string) (TokenList, error) {
	var tokenList TokenList
	path, err := filepath.Abs(listFile)
	if err != nil {
		return tokenList, err
	}
	file, err := os.Open(path)
	if err != nil {
		return tokenList, err
	}
	defer func() { _ = file.Close() }()

	bz, err := ioutil.ReadAll(file)
	if err != nil {
		return tokenList, err
	}
	err = json.Unmarshal(bz, &tokenList)
	return tokenList, err
}

func getMainnetTokens(list TokenList) []Token {
	var newTokens []Token
	for _, token := range list.Tokens {
		if token.ChainID == 101 {
			newTokens = append(newTokens, token)
		}
	}
	return newTokens
}
