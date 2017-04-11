package flags

import "flag"

const (
	DefaultSystem    = "ubuntu-core-16-64"
	DefaultExecutors = 4
)

// Options gathers the given parsed flags
type Options struct {
	System    string
	Executors int
}

// Parse analyzes the given flags and return them inside an Options struct
func Parse() *Options {
	var (
		system    = flag.String("system", DefaultSystem, "spread system to execute the test on")
		executors = flag.Int("executors", DefaultExecutors, "number of parallel testflinger executors")
	)
	flag.Parse()

	return &Options{
		System:    *system,
		Executors: *executors,
	}
}
