package flags_test

import (
	"flag"
	"os"
	"testing"

	"github.com/fgimenez/validator/pkg/flags"
)

func TestParseReturnsParsedFlags(t *testing.T) {
	var parsedFlags interface{}
	parsedFlags = flags.Parse()

	if v, ok := parsedFlags.(*flags.Options); !ok {
		t.Errorf("Parse didn't return options: %q", v)
	}
}

func TestParseSetsSystemToFlagValue(t *testing.T) {
	resetFlag()

	os.Args = []string{"", "-system", "my-system"}
	parsedFlags := flags.Parse()

	if parsedFlags.System != "my-system" {
		t.Errorf("system wasn't parsed: %q instead of my-system", parsedFlags.System)
	}
}

func TestParseSetsSystemToDefaultValue(t *testing.T) {
	resetFlag()

	os.Args = []string{""}
	parsedFlags := flags.Parse()

	if parsedFlags.System != flags.DefaultSystem {
		t.Errorf("system wasn't set to default: %q instead of %q", parsedFlags.System, flags.DefaultSystem)
	}
}

func TestParseSetsexecutorsToFlagValue(t *testing.T) {
	resetFlag()

	os.Args = []string{"", "-executors", "4"}
	parsedFlags := flags.Parse()

	if parsedFlags.Executors != 4 {
		t.Errorf("executors wasn't parsed: %q instead of 4", parsedFlags.Executors)
	}
}

func TestParseSetsExecutorsToDefaultValue(t *testing.T) {
	resetFlag()

	os.Args = []string{""}
	parsedFlags := flags.Parse()

	if parsedFlags.Executors != flags.DefaultExecutors {
		t.Errorf("executors wasn't set to default: %q instead of %q", parsedFlags.Executors, flags.DefaultExecutors)
	}
}

// from flag.ResetForTesting
func resetFlag() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
}
