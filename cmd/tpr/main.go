package main

import (
	"log"

	"github.com/fgimenez/validator/pkg/flags"
	"github.com/fgimenez/validator/pkg/runner"
	"github.com/fgimenez/validator/pkg/types"
)

func main() {
	options := flags.Parse()

	deps := &types.RunnerDependencies{
		C:  cli.New(),
		T:  testflinger.New(),
		Sd: systemd.New(),
		Sp: splitter.New(),
	}
	runner := runner.New(deps)

	if err := runner.Run(options); err != nil {
		log.Fatalf(err)
	}
}
