package runner

import (
	"log"
	"os"
	"strings"

	"github.com/fgimenez/validator/pkg/types"
)

var logger = log.New(os.Stdout, "logger: ", log.Ldate|log.Ltime)

type Runner struct {
	splitter    types.Splitter
	testflinger types.Testflinger
	systemd     types.Systemder
	cli         types.Cli
}

func New(deps *types.RunnerDependencies) *Runner {
	return &Runner{
		splitter:    deps.Sp,
		testflinger: deps.T,
		systemd:     deps.Sd,
		cli:         deps.C,
	}
}

func (r *Runner) Run(options *types.Options) ([]string, error) {
	chunks, err := r.splitter.Run(options)
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
		unitName, unitCmd, err := r.systemd.TransientRunCmd(cfgFile)
		if err != nil {
			log.Printf("Error getting systemd-run command for %v: %v", chunk, err)
			return nil, err
		}
		if _, err := r.cli.ExecCommand(strings.Fields(unitCmd)...); err != nil {
			log.Printf("Error running systemd unit %v with command %v: %v", unitName, unitCmd, err)
			return nil, err
		}
		output = append(output, unitName)
	}
	return output, nil
}
