package main

import (
	"os"

	"github.com/forbole/soljuno/cmd"

	cmdtypes "github.com/forbole/soljuno/cmd/types"
)

func main() {
	// ParsingConfig the runner
	config := cmdtypes.NewConfig("soljuno")

	// Run the commands and panic on any error
	exec := cmd.BuildDefaultExecutor(config)
	err := exec.Execute()
	if err != nil {
		os.Exit(1)
	}
}
