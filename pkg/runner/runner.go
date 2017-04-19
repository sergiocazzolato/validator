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

	chunks, err := r.Splitter.Split(options, strings.Split(list, "\n"))
	if err != nil {
		log.Printf("Error splitting suite: %v", err)
		return nil, err
	}

	var output []string
	for _, chunk := range chunks {
		cfgFile, err := r.Testflinger.GenerateCfg(options, chunk)
		if err != nil {
			log.Printf("Error generating testflinger config for chunk %v: %v", chunk, err)
			return nil, err
		}
		output = append(output, cfgFile)
	}
	return output, nil
}
