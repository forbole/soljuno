package db

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/forbole/soljuno/db"
	"github.com/spf13/cobra"
)

func InitDatabaseCmd(cmdCfg *Config) *cobra.Command {
	return &cobra.Command{
		Use:     "init-db [schema-folder]",
		Short:   "Init the database by schemas in the given folder",
		Args:    cobra.ExactArgs(1),
		PreRunE: ReadConfig(cmdCfg),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := GetDatabaseContext(cmdCfg)
			if err != nil {
				return err
			}
			if err != nil {
				return err
			}
			ctx.Logger.Info("Initializing database...")
			err = InitDatabase(ctx.Database, args[0])
			if err != nil {
				return err
			}
			ctx.Logger.Info("Initialized database...")
			return nil
		},
	}
}

func InitDatabase(db db.ExceutorDb, schemaDir string) error {
	_, err := db.Exec(`DROP SCHEMA public CASCADE;`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE SCHEMA public;`)
	if err != nil {
		return err
	}
	dirPath, err := filepath.Abs(schemaDir)
	if err != nil {
		return err
	}

	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, fileInfo := range dir {
		file, err := ioutil.ReadFile(filepath.Join(dirPath, fileInfo.Name()))
		if err != nil {
			return err
		}

		commentsRegExp := regexp.MustCompile(`/\*.*\*/`)
		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err := db.Exec(commentsRegExp.ReplaceAllString(request, ""))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
