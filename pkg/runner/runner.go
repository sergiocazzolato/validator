package runner

import (
	"log"
	"os"
	"strings"

	"github.com/fgimenez/validator/pkg/types"
)

var logger = log.New(os.Stdout, "logger: ", log.Ldate|log.Ltime)

type Runner struct {
	Splitter    types.Splitter
	Testflinger types.Testflinger
	Cli         types.Cli
}

func New(deps *types.RunnerDependencies) *Runner {
	return &Runner{
		Splitter:    deps.Splitter,
		Testflinger: deps.Testflinger,
		Cli:         deps.Cli,
	}
}

func (r *Runner) Run(options *types.Options) ([]string, error) {
	list, err := r.Cli.ExecCommand("spread", "-list", options.System)
	if err != nil {
		log.Printf("Error getting list: %v", err)
		return nil, err
	}

	chunks := r.Splitter.Split(options, strings.Split(list, "\n"))

	output := r.Testflinger.GenerateCfg(options, chunks)

	return output, nil
}
