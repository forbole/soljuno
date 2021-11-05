package snapshot

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func ImportSnapshotCmd(cmdCfg *Config) *cobra.Command {
	return &cobra.Command{
		Use:     "import [snapshot-file]",
		Short:   "Import a snapshot",
		PreRunE: ReadConfig(cmdCfg),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			context, err := GetParsingContext(cmdCfg)
			if err != nil {
				return err
			}

			return StartImportSnapshot(context, args[0])
		},
	}
}

func StartImportSnapshot(context *Context, snapshotFile string) error {
	ex, err := os.Executable()
	if err != nil {
		return err
	}
	path := filepath.Join(filepath.Dir(ex), snapshotFile)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() { err = file.Close() }()
	reader := bufio.NewReader(file)

	return handle(reader)
}

func handle(reader *bufio.Reader) error {
	_, _, err := reader.ReadLine()
	if err != nil {
		return err
	}
	for {
		account, err := readSection(reader)
		if err == io.EOF {
			break
		}
		fmt.Println(account)
	}
	return nil
}

func readSection(reader *bufio.Reader) (AccountInfo, error) {
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
				return AccountInfo{}, err
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
			return AccountInfo{}, err
		}
	}
	if err == io.EOF {
		return AccountInfo{}, err
	}

	var account AccountInfo
	// Process the line here.
	err = yaml.Unmarshal(detailBuf.Bytes(), &account)
	if err != nil {
		return AccountInfo{}, err
	}
	account.Pubkey = pubkey
	return account, err
}
