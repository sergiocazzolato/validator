package runner

import (
	"log"
	"os"

	"github.com/fgimenez/validator/pkg/types"
)

var logger = log.New(os.Stdout, "logger: ", log.Ldate|log.Ltime)

type Runner struct {
	splitter    types.Splitter
	testflinger types.Testflinger
	cli         types.Cli
}

func New(deps *types.RunnerDependencies) *Runner {
	return &Runner{
		splitter:    deps.Sp,
		testflinger: deps.T,
		cli:         deps.C,
	}
}

func (r *Runner) Run(options *types.Options) ([]string, error) {
	chunks, err := r.splitter.Split(options)
	if err != nil {
		log.Printf("Error splitting suite: %v", err)
		return nil, err
	}

	var output []string
	for _, chunk := range chunks {
		cfgFile, err := r.testflinger.GenerateCfg(options, chunk)
		if err != nil {
			log.Printf("Error generating testflinger config for chunk %v: %v", chunk, err)
			return nil, err
		}
		output = append(output, cfgFile)
	}
	return output, nil
}
