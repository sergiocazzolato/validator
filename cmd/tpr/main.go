package main

import (
	"fmt"
	"log"

	"github.com/fgimenez/validator/pkg/cli"
	"github.com/fgimenez/validator/pkg/flags"
	"github.com/fgimenez/validator/pkg/runner"
	"github.com/fgimenez/validator/pkg/splitter"
	"github.com/fgimenez/validator/pkg/testflinger"
	"github.com/fgimenez/validator/pkg/types"
)

func main() {
	options := flags.Parse()

	deps := &types.RunnerDependencies{
		Cli:         &cli.Executor{},
		Testflinger: &testflinger.Testflinger{},
		Splitter:    &splitter.Splitter{},
	}
	runner := runner.New(deps)

	list, err := runner.Run(options)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(list)
}

