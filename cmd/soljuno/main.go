package main

import (
	"os"

	"github.com/forbole/soljuno/cmd/parse"

	"github.com/forbole/soljuno/modules/registrar"

	"github.com/forbole/soljuno/cmd"
)

func main() {
	// ParsingConfig the runner
	config := cmd.NewConfig("soljuno").
		WithParseConfig(parse.NewConfig().
			WithRegistrar(registrar.NewDefaultRegistrar()),
		)

	// Run the commands and panic on any error
	exec := cmd.BuildDefaultExecutor(config)
	err := exec.Execute()
	if err != nil {
		os.Exit(1)
	}
}
