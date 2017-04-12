package cli

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/fgimenez/validator/pkg/types"
)

const execOutput = "myoutput"

type cliTestSuite struct {
	backExecCommand func(string, ...string) *exec.Cmd
	helperProcess   string
	subject         types.Cli
}

var s = cliTestSuite{}

func TestMain(m *testing.M) {
	s.backExecCommand = execCommand
	execCommand = s.fakeExecCommand
	s.subject = &Executor{}
	s.helperProcess = "TestHelperProcess"

	result := m.Run()
	defer os.Exit(result)

	execCommand = s.backExecCommand
}

func (s *cliTestSuite) fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run", s.helperProcess, "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	baseHelperProcess(0)
}

func TestHelperProcessErr(t *testing.T) {
	baseHelperProcess(1)
}

func baseHelperProcess(exitValue int) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	fmt.Fprintf(os.Stdout, execOutput)
	os.Exit(exitValue)
}

func TestExecCommand(t *testing.T) {
	actualOutput, err := s.subject.ExecCommand("mycmd")
	if err != nil {
		t.Errorf("returned error %v", err)
	}
	if actualOutput != execOutput {
		t.Errorf("expected output %q, obtained %q", execOutput, actualOutput)
	}
}

func TestExecCommandErrWithError(t *testing.T) {
	s.helperProcess = "TestHelperProcessErr"
	defer func() { s.helperProcess = "TestHelperProcess" }()

	actualOutput, err := s.subject.ExecCommand("mycmd")
	if err == nil {
		t.Error("not returned expected error")
	}

	if actualOutput != execOutput {
		t.Errorf("expected output %q, obtained %q", execOutput, actualOutput)
	}
}
