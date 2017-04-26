package types

// Options gathers the given parsed flags
type Options struct {
	System    string
	Executors int
	Channel   string
	From      string
}

// RunnerDependencies entails all the dependencies needed by a runner instance
type RunnerDependencies struct {
	Cli         Cli
	Testflinger Testflinger
	Splitter    Splitter
}

// Cli comprises the methods required by a command manager
type Cli interface {
	ExecCommand(...string) (string, error)
}

// Testflinger represents the methods to interact with the testflinger cli
type Testflinger interface {
	GenerateCfg(*Options, [][]string) []string
}

// Splitter has the methods needed to split the output of spread -list
type Splitter interface {
	Split(*Options, []string) [][]string
}
