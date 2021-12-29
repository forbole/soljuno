package main

import (
	"os"

	"github.com/forbole/soljuno/cmd"
	"github.com/forbole/soljuno/modules/registrar"

	cmdtypes "github.com/forbole/soljuno/cmd/types"
)

func main() {
	// ParsingConfig the runner
	config := cmdtypes.NewConfig("soljuno").WithRegistrar(registrar.NewDefaultRegistrar())

	// Run the commands and panic on any error
	exec := cmd.BuildDefaultExecutor(config)
	err := exec.Execute()
	if err != nil {
		os.Exit(1)
	}
}
