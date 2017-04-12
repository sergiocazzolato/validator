package types

// Options gathers the given parsed flags
type Options struct {
	System    string
	Executors int
}

// RunnerDependencies entails all the dependencies needed by a runner instance
type RunnerDependencies struct {
	C  Cli
	T  Testflinger
	Sd Systemder
	Sp Splitter
}

// Cli comprises the methods required by a command manager
type Cli interface {
	ExecCommand(...string) (string, error)
}

// Testflinger represents the methods to interact with the testflinger cli
type Testflinger interface {
	GenerateCfg(*Options, []string) (string, error)
}

// Systemder comprises the methods related to systemd
type Systemder interface {
	TransientRunCmd(string) (string, string, error)
}

// Splitter has the methods needed to split the output of spread -list
type Splitter interface {
	Run(*Options) ([][]string, error)
}
