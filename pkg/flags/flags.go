package flags

import (
	"flag"

	"github.com/fgimenez/validator/pkg/types"
)

const (
	DefaultSystem    = "ubuntu-core-16-64"
	DefaultExecutors = 4
)

// Parse analyzes the given flags and return them inside an Options struct
func Parse() *types.Options {
	var (
		system    = flag.String("system", DefaultSystem, "spread system to execute the test on")
		executors = flag.Int("executors", DefaultExecutors, "number of parallel testflinger executors")
	)
	flag.Parse()

	return &types.Options{
		System:    *system,
		Executors: *executors,
	}
}
