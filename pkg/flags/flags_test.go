package flags_test

import (
	"flag"
	"os"
	"testing"

	"github.com/fgimenez/validator/pkg/flags"
	"github.com/fgimenez/validator/pkg/types"
)

func TestParseReturnsParsedFlags(t *testing.T) {
	var parsedFlags interface{}
	parsedFlags = flags.Parse()

	if v, ok := parsedFlags.(*types.Options); !ok {
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

func TestParseSetsChannelToFlagValue(t *testing.T) {
	resetFlag()

	os.Args = []string{"", "-channel", "mychannel"}
	parsedFlags := flags.Parse()

	if parsedFlags.Channel != "mychannel" {
		t.Errorf("channel wasn't parsed: %q instead of mychannel", parsedFlags.Channel)
	}
}

func TestParseSetsChannelToDefaultValue(t *testing.T) {
	resetFlag()

	os.Args = []string{""}
	parsedFlags := flags.Parse()

	if parsedFlags.Channel != flags.DefaultChannel {
		t.Errorf("channel wasn't set to default: %q instead of %q", parsedFlags.Channel, flags.DefaultChannel)
	}
}

func TestParseSetsFromToFlagValue(t *testing.T) {
	resetFlag()

	os.Args = []string{"", "-from", "myfrom"}
	parsedFlags := flags.Parse()

	if parsedFlags.From != "myfrom" {
		t.Errorf("from wasn't parsed: %q instead of myfrom", parsedFlags.From)
	}
}

func TestParseSetsFromToDefaultValue(t *testing.T) {
	resetFlag()

	os.Args = []string{""}
	parsedFlags := flags.Parse()

	if parsedFlags.From != flags.DefaultFrom {
		t.Errorf("from wasn't set to default: %q instead of %q", parsedFlags.From, flags.DefaultFrom)
	}
}

func TestParseSetsReleaseToFlagValue(t *testing.T) {
	resetFlag()

	os.Args = []string{"", "-release", "myrelease"}
	parsedFlags := flags.Parse()

	if parsedFlags.Release != "myrelease" {
		t.Errorf("release wasn't parsed: %q instead of myrelease", parsedFlags.Release)
	}
}

func TestParseSetsReleaseToDefaultValue(t *testing.T) {
	resetFlag()

	os.Args = []string{""}
	parsedFlags := flags.Parse()

	if parsedFlags.Release != flags.DefaultRelease {
		t.Errorf("release wasn't set to default: %q instead of %q", parsedFlags.Release, flags.DefaultRelease)
	}
}

// from flag.ResetForTesting
func resetFlag() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
}
