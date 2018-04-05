package flags

import (
	"flag"

	"github.com/fgimenez/validator/pkg/types"
)

const (
	DefaultSystem    = "external:ubuntu-core-16-arm-64"
	DefaultExecutors = 4
	DefaultChannel   = "edge"
	DefaultFrom      = "target"
	DefaultRelease   = "master"
	DefaultQueue     = "dragonboard"
)

// Parse analyzes the given flags and return them inside an Options struct
func Parse() *types.Options {
	var (
		system    = flag.String("system", DefaultSystem, "spread system to execute the test on")
		executors = flag.Int("executors", DefaultExecutors, "number of parallel testflinger executors")
		channel   = flag.String("channel", DefaultChannel, "channel of the target snap to test")
		from      = flag.String("from", DefaultFrom, "determines the channel from which initially provision the image, the target or stable")
		release   = flag.String("release", DefaultRelease, "release branch")
		queue     = flag.String("queue", DefaultQueue, "testflinger queue")
	)
	flag.Parse()

	return &types.Options{
		System:    *system,
		Executors: *executors,
		Channel:   *channel,
		From:      *from,
		Release:   *release,
		Queue:     *queue,
	}
}
